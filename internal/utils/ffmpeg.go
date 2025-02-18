package utils

import (
	"context"
	"fmt"
	"github.com/marcellowy/go-common/gogf/vconfig"
	"github.com/marcellowy/go-common/gogf/vlog"
	"github.com/marcellowy/go-common/os/vstrings"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// CaptureVideoScreenV2 video screenshots
// 上传一张图,是多张图拼在一起的
func CaptureVideoScreenV2(ctx context.Context, videoPath, outputDir string) (pics []string, err error) {
	//picMaxNum := vconfig.Get("tools.maxPicNum", 5).Int()
	var output string
	shell := fmt.Sprintf("ffprobe -v error -show_entries format=duration -of default=noprint_wrappers=1:nokey=1 -i \"%s\"", videoPath)
	if output, err = Exec(ctx, shell); err != nil {
		vlog.Info(ctx, shell)
		vlog.Error(ctx, err)
		return
	}
	output = vstrings.ReplaceLineBreaks(output)

	outputFile := filepath.Join(outputDir, fmt.Sprintf("%%d.jpg"))

	shell = fmt.Sprintf("ffmpeg -y -i \"%s\" -vf \"fps=100/%s,scale=960:-1,tile=2x2\" -frames:v 1 -an \"%s\"", videoPath, output, outputFile)
	if _, err = Exec(ctx, shell); err != nil {
		vlog.Info(ctx, shell)
		vlog.Error(ctx, err)
		return
	}

	pics = append(pics, fmt.Sprintf(outputFile, 1)) // 注意:这里的 1 是 -frames:v 参数的 1,表示截图张数

	return
}

// CaptureVideoScreen video screenshots
// 上传多张图方案
func CaptureVideoScreen(ctx context.Context, mi *MediaInfo, videoPath, outputDir string) (pics []string, err error) {
	rand.Seed(time.Now().UnixNano())
	//vlog.Info(ctx, videoPath, outputDir)
	initialTime := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	picMaxNum := vconfig.Get("tools.maxPicNum", 5).Int()
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < picMaxNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var pic string
			if pic, err = captureVideoScreen(ctx, i, initialTime, mi, videoPath, outputDir); err != nil {
				// 截图出错了
				vlog.Warning(ctx, err)
			}
			if pic != "" {
				mu.Lock()
				pics = append(pics, pic)
				mu.Unlock()
			}
		}()

	}
	wg.Wait()
	return
}

func captureVideoScreen(ctx context.Context, i int, initialTime time.Time, mi *MediaInfo, videoPath, outputDir string) (pic string, err error) {
	n := rand.Intn(60) + 30
	currentTime := initialTime.Add(time.Duration(i) * time.Duration(n) * time.Second)
	// ffmpeg -i input.mp4 -vf "select=eq(n\,100)" -vframes 1 output.jpg
	// -q:v 100
	outputFile := filepath.Join(outputDir, fmt.Sprintf("%d.png", i))    // 普通视频jpg太差,用png
	outputFileJPG := filepath.Join(outputDir, fmt.Sprintf("%d.jpg", i)) // 4k视频png太大,用jpg

	shell := fmt.Sprintf("%s -ss %s -i \"%s\" -q:v 1 -vframes 1 -vsync vfr \"%s\"", "ffmpeg", currentTime.Format("15:04:05"), videoPath, outputFile)
	if mi.IsUHD() {
		shell = fmt.Sprintf("%s -ss %s -i \"%s\" -q:v 2 -vframes 1 -vsync vfr -s 1920x1080 \"%s\"", "ffmpeg", currentTime.Format("15:04:05"), videoPath, outputFile)
	}
	// shell := fmt.Sprintf("ffmpeg -i %s -vf \"select=eq(n\\,%d)\" -vframes 1 \"%s/%d.jpg\"", videoPath, i*10+600, outputDir, i)
	if _, err = Exec(ctx, shell); err != nil {
		vlog.Info(ctx, shell)
		vlog.Error(ctx, err)
		return
	}
	pic = outputFile

	// 判断图片大小,如果图片还是太大,再转为jpg降低质量以降低大小
	var fi os.FileInfo
	if fi, err = os.Stat(outputFile); err != nil {
		//当视频不够长时,这里也会报错,所以可以屏蔽这个错误
		//vlog.Error(ctx, err)
		vlog.Warning(ctx, err)
		// 重置err
		err = nil // 让后面的流程可以继续跑下去，预期之内的错误
		pic = ""
		return
	}
	// 3Mib
	if fi.Size() < 3*1024*1024 {
		return
	}

	shell = fmt.Sprintf("%s -i \"%s\" \"%s\"", "ffmpeg", outputFile, outputFileJPG)
	if _, err = Exec(ctx, shell); err != nil {
		vlog.Info(ctx, shell)
		vlog.Error(ctx, err)
		return
	}
	pic = outputFileJPG

	// 删除png
	_ = os.RemoveAll(outputFile)
	return
}
