package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/marcellowy/go-common/gogf/vconfig"
	"github.com/marcellowy/go-common/gogf/vlog"
	"os"
	"path/filepath"
	"pt-auto/internal/utils"
	"strings"
)

func init() {
	//setting working directory
	var err error
	if err = os.Chdir(filepath.Dir(os.Args[0])); err != nil {
		fmt.Println(err)
	}
}

func GetTorrentNameAndTitle(ctx context.Context, mi *utils.MediaInfo, name, year, suffix, group string) (title, videoFilename, torrentFilename string, err error) {
	videoCodec := utils.VideoCodecStringMapTitle[mi.VideoCodec]
	definition := utils.DefinitionStringMapTitle[mi.Definition]

	title = name + " " + year + " " + definition + " " + videoCodec + "-" + group
	nameWords := strings.Split(name, " ")
	newName := strings.Join(nameWords, ".")
	// 如果翻译后的名字中有 " 分号,会影响执行,去掉分号
	newName = strings.ReplaceAll(newName, "\"", "")
	videoFilename = newName + "." + year + "." + definition + "." + videoCodec + "-" + group + "." + suffix
	torrentFilename = "[M-TEAM]" + newName + "." + year + "." + definition + "." + videoCodec + "-" + group + ".torrent"

	return
}

// .\pt-auto.exe --video-path tools/test-movie.mp4 --year 2015 --name "TEST MOVIE"
var (
	Main = gcmd.Command{
		Name:     "main",
		Usage:    "main",
		Brief:    "start http server",
		Examples: ".\\pt-auto.exe --video-path \"\" --name \"\"",
		Arguments: []gcmd.Argument{
			{
				Name:  "input",
				Short: "i",
			},
			{
				Name:  "year",
				Short: "y",
			},
			{
				Name:  "name",
				Short: "n",
			},
			{
				Name: "subject",
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			//s := g.Server()
			//s.Group("/", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Bind(
			//		hello.NewV1(),
			//	)
			//})
			//s.Run()

			videoPath := parser.GetOpt("input", "").String()
			year := parser.GetOpt("year", "").String()
			name := parser.GetOpt("name", "").String()
			subject := parser.GetOpt("subject", "").String()

			// 如果名字中有双引号会出错
			name = strings.ReplaceAll(name, "\"", "")

			if videoPath == "" {
				vlog.Error(ctx, "param error")
				return
			}

			var vp os.FileInfo
			if vp, err = os.Stat(videoPath); err != nil {
				vlog.Errorf(ctx, "file %s not found", videoPath)
				return err
			}

			s := Service{
				TC: &utils.TranslateConfig{
					Key:  vconfig.Get("tools.rapidapi.key").String(),
					Host: vconfig.Get("tools.rapidapi.host").String(),
				},
				ProxyEnabled: vconfig.Get("tools.proxy.enabled", false).Bool(),
				ProxyAddr:    vconfig.Get("tools.proxy.addr", "").String(),
			}

			if vp.IsDir() {
				vlog.Info(ctx, "directory mode")
				// 如果是目录
				if err = s.Run(ctx, videoPath); err != nil {
					vlog.Error(ctx, err)
					return
				}

				return
			}
			vlog.Info(ctx, "file mode")

			if err = s.RunOnce(ctx, videoPath, year, name, subject); err != nil {
				vlog.Error(ctx, err)
				return
			}

			return nil
		},
	}
)
