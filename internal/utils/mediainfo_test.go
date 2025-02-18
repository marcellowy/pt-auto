package utils

import (
	"context"
	"testing"
)

func TestMediaInfo_Open(t *testing.T) {
	type fields struct {
		Definition       string
		Code             string
		mediaInfoPath    string
		MediaInfoContent string
		mi               []*mediaInfos
	}
	type args struct {
		ctx   context.Context
		video string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				Definition:       "",
				Code:             "",
				mediaInfoPath:    "",
				MediaInfoContent: "",
				mi:               nil,
			},
			args: args{
				ctx: context.Background(),
				// please modify this on local
				video: "E:\\study\\Auto\\pt-auto\\tools\\1409899-uhd_3840_2160_25fps.mp4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MediaInfo{
				Definition: tt.fields.Definition,
				VideoCodec: tt.fields.Code,
				//mediaInfoPath:    tt.fields.mediaInfoPath,
				MediaInfoContent: tt.fields.MediaInfoContent,
				mi:               tt.fields.mi,
			}
			if err := m.Open(tt.args.ctx, tt.args.video); (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
			}
			//fmt.Println(m.MediaInfoContent)
		})
	}
}
