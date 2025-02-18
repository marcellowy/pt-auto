package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/marcellowy/go-common/gogf/vlog"
	"github.com/marcellowy/go-common/tools"
	"html"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var SourceStringMapId = map[string]string{
	SourceWebDL:  SourceWebDLId,
	SourceBluray: SourceBlurayId,
	SourceRemux:  SourceRemuxId,
	SourceEncode: SourceEncodeId,
	SourceHDTV:   SourceHDTVId,
	SourceDVD:    SourceDVDId,
	SourceCD:     SourceCDId,
	SourceOther:  SourceOtherId,
}

const (
	SourceWebDL    = "Web-DL"
	SourceWebDLId  = "8"
	SourceBluray   = "Bluray"
	SourceBlurayId = "1"
	SourceRemux    = "Remux"
	SourceRemuxId  = "4"
	SourceEncode   = "Encode"
	SourceEncodeId = "9"
	SourceHDTV     = "HDTV/TV"
	SourceHDTVId   = "5"
	SourceDVD      = "DVD"
	SourceDVDId    = "3"
	SourceCD       = "CD"
	SourceCDId     = "7"
	SourceOther    = "Other"
	SourceOtherId  = "6"
)

var TeamIdMapString = map[string]string{
	TeamTPTVId:  TeamTPTV,
	TeamMWebId:  TeamMWeb,
	TeamMTeamId: TeamMTeam,
}

const (
	TeamTPTV    = "TPTV"
	TeamTPTVId  = "43"
	TeamMWeb    = "MWeb"
	TeamMWebId  = "44"
	TeamMTeam   = "MTeam"
	TeamMTeamId = "9"
)

// TeamList
//9 MTeam
//23 TnP
//44 MWeb
//43 TPTV
//6 BMDru
//19 CNHK
//8 Pack
//25 CatEDU
//26 ARiC
//27 Telesto
//30 7³ACG
//34 QHstudIo
//31 JKCT
//35 G00DB0Y
//36 D0
//37 iFree
//38 HZH
//50 YzYY
//40 HBO
//41 REE
//42 Bulgur
//45 CTRL
//51 FatTiger
//47 TensoRaws
//48 ZTR
//49 126811

// Category List
// {"message":"SUCCESS","data":{"waterfall":["410","401","419","420","421","439","402","403","435","438","408","434","424","431","437","426","429","430","432","436","440","404","405","406","407","409","411","412","413","422","423","425","427","433","441","442","448"],"adult":["410","429","424","430","426","437","431","432","436","425","433","411","412","413","440"],"tvshow":["403","402","435","438"],"list":[{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 16:52:09","id":"434","order":"1","nameChs":"Music(无损)","nameCht":"Music(無損)","nameEng":"Music(Lossless)","image":"flac.png","parent":"110"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-14 20:58:07","id":"427","order":"1","nameChs":"教育書面","nameCht":"教育(書面)","nameEng":"education book","image":"Study.png","parent":"443"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-26 23:50:33","id":"100","order":"1","nameChs":"电影","nameCht":"電影","nameEng":"Movie","image":"","parent":null},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"401","order":"1","nameChs":"电影/SD","nameCht":"電影/SD","nameEng":"Movie/SD","image":"moviesd.png","parent":"100"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-05-22 13:05:21","id":"423","order":"1","nameChs":"PC游戏","nameCht":"PC遊戲","nameEng":"PCGame","image":"game-pc-3.jpeg","parent":"447"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 16:50:37","id":"403","order":"1","nameChs":"影剧/综艺/SD","nameCht":"影劇/綜藝/SD","nameEng":"TV Series/SD","image":"tvsd.png","parent":"105"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 16:52:05","id":"404","order":"1","nameChs":"纪录","nameCht":"紀錄","nameEng":"Record","image":"bbc.png","parent":"444"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 17:25:00","id":"405","order":"1","nameChs":"动画","nameCht":"動畫","nameEng":"Anime","image":"anime.png","parent":"449"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 17:25:04","id":"407","order":"1","nameChs":"运动","nameCht":"運動","nameEng":"Sports","image":"sport.png","parent":"450"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-26 23:50:36","id":"105","order":"2","nameChs":"影剧/综艺","nameCht":"影劇/綜藝","nameEng":"TV Series","image":"","parent":null},{"createdDate":"2024-04-13 02:02:28","lastModifiedDate":"2024-04-14 20:58:24","id":"441","order":"2","nameChs":"教育(影片)","nameCht":"教育(影片)","nameEng":"edu Video","image":"Study_Video.png","parent":"443"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 17:25:06","id":"422","order":"2","nameChs":"软件","nameCht":"軟體","nameEng":"Software","image":"software.png","parent":"450"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"419","order":"2","nameChs":"电影/HD","nameCht":"電影/HD","nameEng":"Movie/HD","image":"moviehd.png","parent":"100"},{"createdDate":"2024-04-13 17:16:22","lastModifiedDate":"2024-04-13 17:16:31","id":"448","order":"2","nameChs":"TV遊戲","nameCht":"TV遊戲","nameEng":"TvGame","image":"pcgame.png","parent":"447"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 16:50:42","id":"402","order":"2","nameChs":"影剧/综艺/HD","nameCht":"影劇/綜藝/HD","nameEng":"TV Series/HD","image":"tvhd.png","parent":"105"},{"createdDate":"2024-04-13 16:40:33","lastModifiedDate":"2024-04-13 16:40:33","id":"444","order":"3","nameChs":"紀錄","nameCht":"紀錄","nameEng":"BBC","image":null,"parent":null},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"420","order":"3","nameChs":"电影/DVDiSo","nameCht":"電影/DVDiSo","nameEng":"Movie/DVDiSo","image":"moviedvd.png","parent":"100"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 16:52:15","id":"406","order":"3","nameChs":"演唱","nameCht":"演唱","nameEng":"MV","image":"mv.png","parent":"110"},{"createdDate":"2024-04-13 02:03:17","lastModifiedDate":"2024-06-15 02:26:21","id":"442","order":"3","nameChs":"有聲書","nameCht":"有聲書","nameEng":"AuiBook","image":"Study_Audio.png","parent":"450"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 16:50:45","id":"438","order":"3","nameChs":"影剧/综艺/BD","nameCht":"影劇/綜藝/BD","nameEng":"TV Series/BD","image":"tvbd.png","parent":"105"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 17:25:08","id":"409","order":"4","nameChs":"Misc(其他)","nameCht":"Misc(其他)","nameEng":"Misc(Other)","image":"other.png","parent":"450"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"421","order":"4","nameChs":"电影/Blu-Ray","nameCht":"電影/Blu-Ray","nameEng":"Movie/Blu-Ray","image":"moviebd.png","parent":"100"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-26 23:50:49","id":"110","order":"4","nameChs":"Music","nameCht":"Music","nameEng":"Music","image":"","parent":null},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-04-13 16:50:26","id":"435","order":"4","nameChs":"影剧/综艺/DVDiSo","nameCht":"影劇/綜藝/DVDiSo","nameEng":"TV Series/DVDiSo","image":"tvdvd.png","parent":"105"},{"createdDate":"2024-04-13 15:02:04","lastModifiedDate":"2024-04-13 15:02:13","id":"443","order":"5","nameChs":"教育","nameCht":"教育","nameEng":"eu","image":null,"parent":null},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"439","order":"5","nameChs":"电影/Remux","nameCht":"電影/Remux","nameEng":"Movie/Remux","image":"movieremux.png","parent":"100"},{"createdDate":"2024-04-13 17:15:28","lastModifiedDate":"2024-04-13 17:15:37","id":"447","order":"6","nameChs":"遊戲","nameCht":"遊戲","nameEng":"遊戲","image":null,"parent":null},{"createdDate":"2024-04-13 17:22:46","lastModifiedDate":"2024-04-13 17:22:55","id":"449","order":"7","nameChs":"動漫","nameCht":"動漫","nameEng":"Anime","image":null,"parent":null},{"createdDate":"2024-04-13 17:24:09","lastModifiedDate":"2024-04-13 17:24:09","id":"450","order":"8","nameChs":"其他","nameCht":"其他","nameEng":"其他","image":null,"parent":null},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-26 23:51:46","id":"115","order":"20","nameChs":"AV(有码)","nameCht":"AV(有碼)","nameEng":"AV(有碼)","image":"","parent":null},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-26 23:51:50","id":"120","order":"21","nameChs":"AV(无码)","nameCht":"AV(無碼)","nameEng":"AV(無碼)","image":"","parent":null},{"createdDate":"2024-04-13 16:52:43","lastModifiedDate":"2024-04-13 16:52:51","id":"445","order":"22","nameChs":"IV","nameCht":"IV","nameEng":"IV","image":null,"parent":null},{"createdDate":"2024-04-13 16:53:44","lastModifiedDate":"2024-04-13 16:53:44","id":"446","order":"23","nameChs":"H-ACG","nameCht":"H-ACG","nameEng":"H-ACG","image":null,"parent":null},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"410","order":"31","nameChs":"AV(有码)/HD Censored","nameCht":"AV(有碼)/HD Censored","nameEng":"AV(有碼)/HD Censored","image":"cenhd.png","parent":"115"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"429","order":"32","nameChs":"AV(无码)/HD Uncensored","nameCht":"AV(無碼)/HD Uncensored","nameEng":"AV(無碼)/HD Uncensored","image":"uenhd.png","parent":"120"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"424","order":"33","nameChs":"AV(有码)/SD Censored","nameCht":"AV(有碼)/SD Censored","nameEng":"AV(有碼)/SD Censored","image":"censd.png","parent":"115"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"430","order":"34","nameChs":"AV(无码)/SD Uncensored","nameCht":"AV(無碼)/SD Uncensored","nameEng":"AV(無碼)/SD Uncensored","image":"uensd.png","parent":"120"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"426","order":"35","nameChs":"AV(无码)/DVDiSo Uncensored","nameCht":"AV(無碼)/DVDiSo Uncensored","nameEng":"AV(無碼)/DVDiSo Uncensored","image":"uendvd.png","parent":"120"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"437","order":"36","nameChs":"AV(有码)/DVDiSo Censored","nameCht":"AV(有碼)/DVDiSo Censored","nameEng":"AV(有碼)/DVDiSo Censored","image":"cendvd.png","parent":"115"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"431","order":"37","nameChs":"AV(有码)/Blu-Ray Censored","nameCht":"AV(有碼)/Blu-Ray Censored","nameEng":"AV(有碼)/Blu-Ray Censored","image":"cenbd.png","parent":"115"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"432","order":"38","nameChs":"AV(无码)/Blu-Ray Uncensored","nameCht":"AV(無碼)/Blu-Ray Uncensored","nameEng":"AV(無碼)/Blu-Ray Uncensored","image":"uenbd.png","parent":"120"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"436","order":"39","nameChs":"AV(网站)/0Day","nameCht":"AV(網站)/0Day","nameEng":"AV(網站)/0Day","image":"adult0day.png","parent":"120"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"425","order":"40","nameChs":"IV(写真影集)","nameCht":"IV(寫真影集)","nameEng":"IV/Video Collection","image":"ivvideo.png","parent":"445"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"433","order":"41","nameChs":"IV(写真图集)","nameCht":"IV(寫真圖集)","nameEng":"IV/Picture Collection","image":"ivpic.png","parent":"445"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"411","order":"51","nameChs":"H-游戏","nameCht":"H-遊戲","nameEng":"H-Game","image":"hgame.png","parent":"446"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"412","order":"52","nameChs":"H-动漫","nameCht":"H-動畫","nameEng":"H-Anime","image":"hanime.png","parent":"446"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"413","order":"53","nameChs":"H-漫画","nameCht":"H-漫畫","nameEng":"H-Comic","image":"hcomic.png","parent":"446"},{"createdDate":"2024-03-22 14:00:15","lastModifiedDate":"2024-03-22 14:00:15","id":"440","order":"440","nameChs":"AV(Gay)/HD","nameCht":"AV(Gay)/HD","nameEng":"AV(Gay)/HD","image":"gayhd.gif","parent":"120"}],"music":["406","434"],"movie":["401","419","420","421","439"]},"code":"0"}
type MTeamApi struct {
	URL             string
	ApiKey          string
	UploadImgApiKey string
	// 以下表单部分内容
	SourceId   string
	TeamId     string
	CategoryId string
	// 网络代理
	ProxyEnabled bool
	ProxyAddr    string
}

//proxyEnabled := vconfig.Get("tools.proxy.enabled", false).Bool()
//	proxyAddr := vconfig.Get("tools.proxy.addr", "").String()

type CategoryListData struct {
	CreatedDate      string `json:"createdDate"`
	LastModifiedDate string `json:"lastModifiedDate"`
	ID               string `json:"id"`
	Order            string `json:"order"`
	NameChs          string `json:"nameChs"`
	NameCht          string `json:"nameCht"`
	NameEng          string `json:"nameEng"`
	Image            string `json:"image"`
	Parent           string `json:"parent"`
}

type CategoryListResponse struct {
	Message string `json:"message"`
	Data    struct {
		List []CategoryListData `json:"list"`
	} `json:"data"`
	Code string `json:"code"`
}

func (m *MTeamApi) CategoryList(ctx context.Context) (resp *CategoryListResponse, err error) {
	path := "/api/torrent/categoryList"
	data := make(map[string]string)
	var body []byte
	if body, err = m.Send(ctx, path, data); err != nil {
		vlog.Error(ctx, err)
		return
	}
	var res = CategoryListResponse{}
	if err = json.Unmarshal(body, &res); err != nil {
		vlog.Error(ctx, err)
		return
	}
	//fmt.Println(res.Data.List[0].NameChs)
	resp = &res

	return
}

type AudioCodecListResponse struct {
	Message string `json:"message"`
	Data    []struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Order string `json:"order"`
	} `json:"data"`
	Code string `json:"code"`
}

func (m *MTeamApi) AudioCodecList(ctx context.Context) (resp *AudioCodecListResponse, err error) {
	path := "/api/torrent/audioCodecList"
	//path := "/api/torrent/mediumList"
	data := make(map[string]string)
	var body []byte
	if body, err = m.Send(ctx, path, data); err != nil {
		vlog.Error(ctx, err)
		return
	}
	//fmt.Println(string(body))
	var res = AudioCodecListResponse{}
	if err = json.Unmarshal(body, &res); err != nil {
		vlog.Error(ctx, err)
		return
	}
	resp = &res

	return
}

type MediumListResponse struct {
	Message string `json:"message"`
	Data    []struct {
		ID      string `json:"id"`
		Order   string `json:"order"`
		NameChs string `json:"nameChs"`
		NameCht string `json:"nameCht"`
		NameEng string `json:"nameEng"`
	} `json:"data"`
	Code string `json:"code"`
}

func (m *MTeamApi) MediumList(ctx context.Context) (resp *MediumListResponse, err error) {
	path := "/api/torrent/mediumList"
	data := make(map[string]string)
	var body []byte
	if body, err = m.Send(ctx, path, data); err != nil {
		vlog.Error(ctx, err)
		return
	}

	var res = MediumListResponse{}
	if err = json.Unmarshal(body, &res); err != nil {
		vlog.Error(ctx, err)
		return
	}
	resp = &res

	return
}

func (m *MTeamApi) SourceList(ctx context.Context) (resp *MediumListResponse, err error) {
	path := "/api/torrent/sourceList"
	data := make(map[string]string)
	var body []byte
	if body, err = m.Send(ctx, path, data); err != nil {
		vlog.Error(ctx, err)
		return
	}
	//fmt.Println(string(body))
	var res = MediumListResponse{}
	if err = json.Unmarshal(body, &res); err != nil {
		vlog.Error(ctx, err)
		return
	}
	resp = &res

	return
}

type VideoCodecListResponse struct {
	Message string `json:"message"`
	Data    []struct {
		ID    string `json:"id"`
		Order string `json:"order"`
		Name  string `json:"name"`
	} `json:"data"`
	Code string `json:"code"`
}

func (m *MTeamApi) VideoCodecList(ctx context.Context) (resp *VideoCodecListResponse, err error) {
	path := "/api/torrent/videoCodecList"
	data := make(map[string]string)
	var body []byte
	if body, err = m.Send(ctx, path, data); err != nil {
		vlog.Error(ctx, err)
		return
	}
	//fmt.Println(string(body))
	var res = VideoCodecListResponse{}
	if err = json.Unmarshal(body, &res); err != nil {
		vlog.Error(ctx, err)
		return
	}
	resp = &res

	return
}

type TeamListResponse struct {
	Message string `json:"message"`
	Data    []struct {
		ID     string `json:"id"`
		Order  string `json:"order"`
		Name   string `json:"name"`
		Leader string `json:"leader"`
	} `json:"data"`
	Code string `json:"code"`
}

func (m *MTeamApi) TeamList(ctx context.Context) (resp *TeamListResponse, err error) {
	path := "/api/torrent/teamList"
	data := make(map[string]string)
	var body []byte
	if body, err = m.Send(ctx, path, data); err != nil {
		vlog.Error(ctx, err)
		return
	}

	var res = TeamListResponse{}
	if err = json.Unmarshal(body, &res); err != nil {
		vlog.Error(ctx, err)
		return
	}
	resp = &res
	return
}

type StandardListResponse struct {
	Message string `json:"message"`
	Data    []struct {
		ID    string `json:"id"`
		Order string `json:"order"`
		Name  string `json:"name"`
	} `json:"data"`
	Code string `json:"code"`
}

func (m *MTeamApi) StandardList(ctx context.Context) (resp *StandardListResponse, err error) {
	path := "/api/torrent/standardList"
	data := make(map[string]string)
	var body []byte
	if body, err = m.Send(ctx, path, data); err != nil {
		vlog.Error(ctx, err)
		return
	}
	fmt.Println(string(body))
	var res = StandardListResponse{}
	if err = json.Unmarshal(body, &res); err != nil {
		vlog.Error(ctx, err)
		return
	}
	resp = &res
	return
}

func (m *MTeamApi) Send(ctx context.Context, path string, data map[string]string) (body []byte, err error) {

	client := gclient.New()
	client.SetHeader("x-api-key", m.ApiKey)
	// 如果x-api-key错误,将会提示: {"message":"本地時間誤差過大，請校準系統時間","data":null,"code":1}
	if err = SetProxy(ctx, client, m.ProxyEnabled, m.ProxyAddr); err != nil {
		vlog.Error(ctx, err)
		return
	}

	var response *gclient.Response
	response, err = client.Post(ctx, m.URL+path, data)
	if err != nil {
		vlog.Error(ctx, err)
		return
	}
	defer tools.Close(response)

	//vlog.Info(ctx, "Status Code: ", response.StatusCode)

	if body, err = io.ReadAll(response.Body); err != nil {
		vlog.Error(ctx, err)
		return
	}
	defer tools.Close(response.Body)

	//vlog.Info(ctx, "response body:", string(body))
	return
}

type sendFormUpload struct {
	FieldName string
	Path      string
}

func (m *MTeamApi) createFormFile(ctx context.Context, fieldName, filename string, writer **multipart.Writer) (err error) {
	// vlog.Info(ctx, "upload file", filename)
	// vlog.Info(ctx, "upload filename", filepath.Base(filename))
	var uploadWriter io.Writer
	if uploadWriter, err = (*writer).CreateFormFile(fieldName, filepath.Base(filename)); err != nil {
		vlog.Error(ctx, err)
		return
	}
	var uploadFile *os.File
	if uploadFile, err = os.Open(filename); err != nil {
		vlog.Error(ctx, err)
		return
	}
	defer tools.Close(uploadFile)
	if _, err = io.Copy(uploadWriter, uploadFile); err != nil {
		vlog.Error(ctx, err)
		return
	}

	return
}

// SendForm 不能用gogf的sendForm
// 当发送的表单中的字段存在 = 号时,会被奇怪的截断
func (m *MTeamApi) SendForm(ctx context.Context, path string, data map[string]string) (body []byte, err error) {

	bb := &bytes.Buffer{}
	writer := multipart.NewWriter(bb)
	var uploadFiles []*sendFormUpload
	for k, v := range data {
		if strings.Contains(v, "@file:") {
			uploadFiles = append(uploadFiles, &sendFormUpload{
				FieldName: k,
				Path:      strings.ReplaceAll(v, "@file:", ""),
			})

			continue
		}
		_ = writer.WriteField(k, v)
	}

	// 上传文件
	for _, v := range uploadFiles {
		if err = m.createFormFile(ctx, v.FieldName, v.Path, &writer); err != nil {
			vlog.Error(ctx, err)
			return
		}
	}
	_ = writer.Close()

	client := http.Client{}
	if m.ProxyEnabled {
		var proxy *url.URL
		if proxy, err = url.Parse(m.ProxyAddr); err != nil {
			vlog.Error(ctx, err, m.ProxyAddr)
			return
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}

	var request *http.Request
	if request, err = http.NewRequest("POST", m.URL+path, bb); err != nil {
		vlog.Error(ctx, err)
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("x-api-key", m.ApiKey)

	var response *http.Response
	if response, err = client.Do(request); err != nil {
		vlog.Error(ctx, err)
		return nil, err
	}
	defer tools.Close(response.Body)

	if response.StatusCode != http.StatusOK {
		vlog.Warning(ctx, "Status Code: ", response.StatusCode)
		return
	}

	if body, err = io.ReadAll(response.Body); err != nil {
		vlog.Error(ctx, err)
		return
	}
	// vlog.Info(ctx, "response body:", string(body))
	return
}

type UploadImgRespData struct {
	StatusCode int `json:"status_code"`
	Image      struct {
		Url string `json:"url"`
	} `json:"image"`
}

func (m *MTeamApi) UploadImg(ctx context.Context, path string) (imgUrl string, err error) {

	vlog.Info(ctx, "upload:", path)

	client := gclient.New()
	client.SetHeader("x-api-key", m.UploadImgApiKey)
	if err = SetProxy(ctx, client, m.ProxyEnabled, m.ProxyAddr); err != nil {
		vlog.Error(ctx, err)
		return
	}

	var response *gclient.Response
	response, err = client.Post(ctx, "https://img.m-team.cc/api/1/upload", "source=@file:"+path)
	if err != nil {
		vlog.Error(ctx, err)
		return
	}
	defer tools.Close(response)

	//vlog.Info(ctx, "Status Code: ", response.StatusCode)

	var body []byte
	if body, err = io.ReadAll(response.Body); err != nil {
		vlog.Error(ctx, err)
		return
	}
	defer tools.Close(response.Body)

	//vlog.Info(ctx, "response body:", string(body))

	var data UploadImgRespData
	if err = json.Unmarshal(body, &data); err != nil {
		vlog.Error(ctx, err)
		return "", err
	}

	imgUrl = data.Image.Url
	//vlog.Info(ctx, "url:", imgUrl)

	return
}

func (m *MTeamApi) GetImgHTML(url string) string {
	return fmt.Sprintf(`![](%s)`, url)
}

type PostForm struct {
	File             string   // 种子文件
	CategoryId       string   // 类别
	Name             string   // 标题
	SmallDescription string   // 副标题
	SourceId         string   // 来源
	StandardId       string   // 解析度
	VideoCodecId     string   // 视频编码
	AudioCodecId     string   // 音频编码
	TeamId           string   // 制作组
	Imdb             string   // IMDB链接
	DouBan           string   // 豆瓣链接
	LabelsNew        []string // 标记
	Mediainfo        string   // MediaInfo 檔案
	Description      string   // 简介
	Anonymous        bool     // 是否匿名
}

type CreateOrEditResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (m *MTeamApi) CreateOrEdit(ctx context.Context, pf *PostForm) (err error) {
	path := "/api/torrent/createOredit"
	data := make(map[string]string)
	{
		data["file"] = "@file:" + pf.File
		data["category"] = pf.CategoryId
		data["name"] = pf.Name
		data["smallDescr"] = pf.SmallDescription
		//data["dmmCode"] = "null" // 不知道是什么
		data["source"] = pf.SourceId
		data["standard"] = pf.StandardId
		data["videoCodec"] = pf.VideoCodecId
		data["audioCodec"] = pf.AudioCodecId
		data["team"] = pf.TeamId
		data["imdb"] = pf.Imdb
		data["douban"] = pf.DouBan
		//data["labelsNew"] = strings.Join(pf.LabelsNew, ",")
		data["mediainfo"] = html.EscapeString(pf.Mediainfo)
		//data["cids"] = ""
		//data["tags"] = ""
		data["anonymous"] = "false"
		if pf.Anonymous {
			data["anonymous"] = "true"
		}
		//data["aids"] = ""
		data["descr"] = pf.Description
		//data["labels"] = "0"
	}

	cc, _ := json.Marshal(data)
	_ = os.WriteFile("tmp.txt", cc, os.ModePerm)

	var body []byte
	if body, err = m.SendForm(ctx, path, data); err != nil {
		vlog.Error(ctx, err)
		return
	}
	var res = CreateOrEditResponse{}
	if err = json.Unmarshal(body, &res); err != nil {
		vlog.Error(ctx, err, string(body))
		return
	}

	if res.Code != "0" {
		err = fmt.Errorf(res.Message)
		vlog.Error(ctx, err)
		return
	}
	return
}
