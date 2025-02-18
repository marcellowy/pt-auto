package utils

import (
	"context"
	"fmt"
	"testing"
)

func TestCaptureVideoScreen(t *testing.T) {
	type args struct {
		ctx       context.Context
		mi        *MediaInfo
		videoPath string
		outputDir string
	}

	var videoPath = "E:\\study\\Auto\\pt-auto\\tools\\1409899-uhd_3840_2160_25fps.mp4"

	mi := &MediaInfo{}
	if err := mi.Open(context.Background(), videoPath); err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		name     string
		args     args
		wantPics []string
		wantErr  bool
	}{
		{
			name: "test",
			args: args{
				ctx:       context.Background(),
				mi:        mi,
				videoPath: videoPath,
				outputDir: "E:\\study\\Auto\\pt-auto\\tools\\",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPics, err := CaptureVideoScreen(tt.args.ctx, tt.args.mi, tt.args.videoPath, tt.args.outputDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("CaptureVideoScreen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotPics) == 0 {
				t.Errorf("CaptureVideoScreen() gotPics = %v, want %v", gotPics, tt.wantPics)
			}
		})
	}
}

func TestCaptureVideoScreenV2(t *testing.T) {
	type args struct {
		ctx       context.Context
		videoPath string
		outputDir string
	}
	tests := []struct {
		name     string
		args     args
		wantPics []string
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				ctx:       context.Background(),
				videoPath: "E:/study/Auto/pt-auto/tools/永夜星河.mp4",
				outputDir: "E:/study/Auto/pt-auto/tools",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPics, err := CaptureVideoScreenV2(tt.args.ctx, tt.args.videoPath, tt.args.outputDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("CaptureVideoScreenV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotPics)
		})
	}
}
