package cmd

import (
	"context"
	"github.com/marcellowy/go-common/gogf/vconfig"
	"github.com/marcellowy/go-common/gogf/vlog"
	"github.com/marcellowy/go-common/tools"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path"
	"path/filepath"
	"pt-auto/internal/utils"
	"strings"
	"sync"
	"time"
)

type Service struct {
	TC *utils.TranslateConfig
	// 网络代理
	ProxyEnabled bool
	ProxyAddr    string
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) readDir(ctx context.Context, dir string, filter func(name string) bool) (filename []string, err error) {
	var entry []os.DirEntry
	entry, err = os.ReadDir(dir)
	if err != nil {
		vlog.Error(ctx, err)
		return
	}
	for _, n := range entry {
		if n.IsDir() {
			continue
		}
		if filter == nil || filter(n.Name()) {
			//vlog.Info(ctx, "add", n.Name())
			filename = append(filename, n.Name())
		}
	}
	return
}

// Run 读取 vDirectory 目录下所有文件并进行处理
func (service *Service) Run(ctx context.Context, vDirectory string) (err error) {
	var filename []string

	if filename, err = service.readDir(ctx, vDirectory, func(name string) bool {
		//vlog.Info(ctx, filepath.Ext(name), dotSuffix)
		return true
	}); err != nil {
		vlog.Error(ctx, err)
		return
	}

	// 解析: name##year##[subject].ts 文件
	for _, fn := range filename {
		vlog.Info(ctx, "filename:", fn)
		var realFilename string
		if realFilename, err = filepath.Abs(filepath.Join(vDirectory, fn)); err != nil {
			vlog.Error(ctx, err)
			continue
		}
		//suffix := strings.ReplaceAll(filepath.Ext(realFilename), ".", "")
		//fn = strings.ReplaceAll(fn, "."+suffix, "")
		//fnArr := strings.Split(fn, "##")
		////vlog.Info(ctx, fnArr)
		//if len(fnArr) < 2 {
		//	vlog.Warningf(ctx, "%s no ##", fn)
		//	continue
		//}
		//nameZH := fnArr[0]
		//year := fnArr[1]
		//
		//subject := ""
		//if len(fnArr) == 3 {
		//	subject = fnArr[2]
		//}
		//
		//// Translate name
		//var nameEN string
		//if nameEN, err = utils.TranslateText(ctx, service.TC, nameZH); err != nil {
		//	vlog.Error(ctx, err)
		//	continue
		//}
		//vlog.Info(ctx, nameZH, "trans:", nameEN)

		if err = service.RunOnce(ctx, realFilename, "", "", ""); err != nil {
			vlog.Error(ctx, err)
		}
	}

	return
}

// input: G:\\游戏娱乐\\街机游戏合集300.zip
// return: 街机游戏合集300
func (service *Service) parseNameFromVideoPath(ctx context.Context, videoPath string) (nameChs string) {
	suffix := strings.ReplaceAll(filepath.Ext(videoPath), ".", "") // 获取扩展名
	filename := filepath.Base(videoPath)                           // 获取文件名部分
	return strings.ReplaceAll(filename, "."+suffix, "")            // 仅留下中文名
}

const (
	NameFlagNameChs = 2 << iota
	NameFlagYear
	NameFlagSubject
	NameFlagNameEng
)

type NameInfo struct {
	NameChs string
	Year    string
	Subject string
	NameEng string
	Flag    int
}

func (service *Service) parseNameByCSharp(ctx context.Context, nameChs string) (info *NameInfo, err error) {
	var (
		symbol = "##"
	)
	info = &NameInfo{}
	if !strings.Contains(nameChs, symbol) {
		info.NameChs = nameChs
		info.Flag |= NameFlagNameChs
	}
	// 如果包含##说明这个文件名是带格式的
	// 如果这一段使用默认值,就用 - 代替
	// nameChs##[Year]##[Subject]##[nameEng].ts 文件
	count := strings.Count(nameChs, symbol)
	//vlog.Info(ctx, "count:", count)
	nameArr := strings.Split(nameChs, symbol)
	for i := 0; i <= count; i++ {
		val := strings.TrimSpace(nameArr[i])
		//vlog.Info(ctx, i, val)
		if val == "-" {
			continue
		}
		switch i {
		case 0:
			info.NameChs = val
			info.Flag |= NameFlagNameChs
		case 1:
			info.Year = val
			info.Flag |= NameFlagYear
		case 2:
			info.Subject = val
			info.Flag |= NameFlagSubject
		case 3:
			info.NameEng = val
			info.Flag |= NameFlagNameEng
		}
	}

	// vlog.Info(ctx, "fill value")
	// 这里补齐没有的项
	if info.Flag&NameFlagYear != NameFlagYear {
		info.Year = time.Now().Format("2006")
		info.Flag |= NameFlagYear
	}

	if info.Flag&NameFlagSubject != NameFlagSubject {
		info.Subject = info.NameChs // 没有定义subject,就使用文件名
		info.Flag |= NameFlagSubject
	}

	if info.Flag&NameFlagNameEng != NameFlagNameEng {
		// 如果没有英文名,就自动翻译一个
		if info.NameEng, err = utils.TranslateText(ctx, service.TC, info.NameChs, service.ProxyEnabled, service.ProxyAddr); err != nil {
			vlog.Error(ctx, err)
			return
		}
		info.Flag |= NameFlagNameEng
	}

	return
}

func (service *Service) RunOnce(ctx context.Context, vPath, year, name, subject string) (err error) {

	if vPath == "" {
		vlog.Error(ctx, "param error")
		return
	}

	nameChs := service.parseNameFromVideoPath(ctx, vPath)
	var nameInfo *NameInfo
	if nameInfo, err = service.parseNameByCSharp(ctx, nameChs); err != nil {
		vlog.Error(ctx, err)
		return
	}

	// 参数传进来有值,就优先使用参数上的值
	if year != "" {
		nameInfo.Year = year
	}

	if name != "" {
		nameInfo.NameEng = name
	}

	if subject != "" {
		nameInfo.Subject = subject
	}

	if _, err = os.Stat(vPath); err != nil {
		vlog.Errorf(ctx, "file %s not found", vPath)
		return err
	}

	// 视频文件后缀
	suffix := strings.ReplaceAll(filepath.Ext(vPath), ".", "")

	// 把单词的首字母都修改为大写
	titleCases := cases.Title(language.English, cases.NoLower)
	name = titleCases.String(nameInfo.NameEng)

	// 输出目录
	outputDirectory := path.Join(filepath.Dir(vPath), strings.ReplaceAll(filepath.Base(vPath), "."+suffix, ""))
	vlog.Info(ctx, "output directory:", outputDirectory)

	// 重置
	_ = os.RemoveAll(outputDirectory)
	_ = os.MkdirAll(outputDirectory, os.ModePerm)

	mi := &utils.MediaInfo{}
	if err = mi.Open(ctx, vPath); err != nil {
		vlog.Error(ctx, err)
		return
	}

	// screenshots
	var pics []string
	if pics, err = utils.CaptureVideoScreenV2(ctx, vPath, outputDirectory); err != nil {
		vlog.Error(ctx, "capture screen failed!", err)
		return err
	}

	// gen name
	//TeamIdMapString
	group := vconfig.Get("tools.mteam.teamName", "").String()
	title, videoFilename, torrentFilename, err := GetTorrentNameAndTitle(ctx, mi, name, nameInfo.Year, suffix, group)
	if err != nil {
		vlog.Error(ctx, "gen name failed!", err)
		return err
	}

	if subject == "" {
		subject = nameInfo.Subject
	}

	//var newVideoPath string
	var movieNewPath = path.Join(outputDirectory, videoFilename) // 默认路径
	if vconfig.Get("tools.torrentMapVideoDirEnabled", false).Bool() {
		dir := vconfig.Get("tools.torrentMapVideoDir").String()
		movieNewPath = path.Join(dir, videoFilename) // 将文件copy到指定路径,方便搬运或者做种
	}

	if err = tools.CopyFile(movieNewPath, vPath); err != nil {
		vlog.Error(ctx, "copy movie failed!", err)
		return err
	}

	var file = path.Join(outputDirectory, torrentFilename)
	if err = utils.CreateTorrent(ctx, movieNewPath, file); err != nil {
		vlog.Error(ctx, file)
		vlog.Error(ctx, "gen torrent failed!")
		return err
	}

	//fmt.Println(title)
	//fmt.Println(path.Join(outputDir, videoFilename))
	//fmt.Println(path.Join(outputDir, torrentFilename))
	//fmt.Println(subject)

	content := "Base:\r\n"
	content += title + "\r\n"
	content += subject + "\r\n"
	content += "\r\n"
	content += mi.MediaInfoContent + "\r\n"

	if err = os.WriteFile(path.Join(outputDirectory, "output.txt"), []byte(content), os.ModePerm); err != nil {
		vlog.Error(ctx, "write file output.txt failed!", err)
		return err
	}

	// 发布
	if err = service.Publish(ctx, file, title, subject, pics, mi); err != nil {
		return
	}

	return
}

func (service *Service) Publish(ctx context.Context, file, name, subSubject string, pics []string, mi *utils.MediaInfo) (err error) {

	api := utils.MTeamApi{
		URL:             vconfig.Get("tools.mteam.URL", "").String(),
		ApiKey:          vconfig.Get("tools.mteam.apiKey", "").String(),
		UploadImgApiKey: vconfig.Get("tools.mteam.uploadImgKey", "").String(),
		SourceId:        vconfig.Get("tools.mteam.sourceId", "").String(),
		TeamId:          vconfig.Get("tools.mteam.teamId", "").String(),
		CategoryId:      vconfig.Get("tools.mteam.categoryId", "").String(), // 影剧/综艺/SD

		ProxyEnabled: vconfig.Get("tools.proxy.enabled", false).Bool(),
		ProxyAddr:    vconfig.Get("tools.proxy.addr", "").String(),
	}

	// 上传图片
	var description []string
	var mu sync.Mutex
	var wg sync.WaitGroup
	// 异步同时上传多张
	for _, v := range pics {
		wg.Add(1)
		go func(picUrl string) {
			defer wg.Done()
			var url string
			if url, err = api.UploadImg(ctx, picUrl); err == nil {
				vlog.Info(ctx, "Upload Image:", picUrl, "Success")
				mu.Lock()
				description = append(description, api.GetImgHTML(url))
				mu.Unlock()
			} else {
				vlog.Info(ctx, "Upload Image:", picUrl, "Failed")
				vlog.Error(ctx, err)
				return
			}
		}(v)
	}
	wg.Wait()

	form := utils.PostForm{
		File:             file,
		CategoryId:       api.CategoryId,
		Name:             name,
		SmallDescription: subSubject,
		SourceId:         api.SourceId,
		StandardId:       utils.DefinitionStringMapId[mi.Definition],
		VideoCodecId:     utils.VideoCodecStringMapId[mi.VideoCodec],
		AudioCodecId:     utils.MediaAudioCodecStringMapId[mi.AudioCodec],
		TeamId:           api.TeamId,
		Imdb:             "",
		DouBan:           "",
		LabelsNew:        nil,
		Mediainfo:        mi.MediaInfoContent,
		Description:      strings.Join(description, " "),
		Anonymous:        true,
	}

	if form.Description == "" {
		// 这个字段为必填字段,当没有图片时添加一个点位符
		form.Description = "placeholder"
	}

	vlog.Info(ctx, "publish Name:", form.Name)
	vlog.Info(ctx, "subTitle:", form.SmallDescription)
	if err = api.CreateOrEdit(ctx, &form); err != nil {
		vlog.Error(ctx, "create or edit failed!", err)
	}
	vlog.Info(ctx, "publish success")
	return
}
