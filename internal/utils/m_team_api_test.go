package utils

import (
	"context"
	"fmt"
	"testing"
)

func TestMTeamApi_AudioCodecList(t *testing.T) {
	type fields struct {
		URL             string
		ApiKey          string
		UploadImgApiKey string
		SourceId        string
		TeamId          string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *AudioCodecListResponse
		wantErr  bool
	}{
		{
			name: "",
			fields: fields{
				URL:             "https://test2.m-team.cc",
				ApiKey:          "",
				UploadImgApiKey: "",
				SourceId:        "",
				TeamId:          "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MTeamApi{
				URL:             tt.fields.URL,
				ApiKey:          tt.fields.ApiKey,
				UploadImgApiKey: tt.fields.UploadImgApiKey,
				SourceId:        tt.fields.SourceId,
				TeamId:          tt.fields.TeamId,
			}
			gotResp, err := m.AudioCodecList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AudioCodecList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotResp.Data {
				fmt.Println(v.Name, v.ID)
			}
		})
	}
}

func TestMTeamApi_SourceList(t *testing.T) {
	type fields struct {
		URL             string
		ApiKey          string
		UploadImgApiKey string
		SourceId        string
		TeamId          string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *MediumListResponse
		wantErr  bool
	}{
		{
			fields: fields{
				URL:             "",
				ApiKey:          "",
				UploadImgApiKey: "",
				SourceId:        "",
				TeamId:          "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MTeamApi{
				URL:             tt.fields.URL,
				ApiKey:          tt.fields.ApiKey,
				UploadImgApiKey: tt.fields.UploadImgApiKey,
				SourceId:        tt.fields.SourceId,
				TeamId:          tt.fields.TeamId,
			}
			gotResp, err := m.SourceList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SourceList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotResp.Data {
				fmt.Println(v.ID, v.NameChs, v.NameCht, v.NameEng)
			}
		})
	}
}

func TestMTeamApi_VideoCodecList(t *testing.T) {
	type fields struct {
		URL             string
		ApiKey          string
		UploadImgApiKey string
		SourceId        string
		TeamId          string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *VideoCodecListResponse
		wantErr  bool
	}{
		{
			fields: fields{
				URL:             "",
				ApiKey:          "",
				UploadImgApiKey: "",
				SourceId:        "",
				TeamId:          "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MTeamApi{
				URL:             tt.fields.URL,
				ApiKey:          tt.fields.ApiKey,
				UploadImgApiKey: tt.fields.UploadImgApiKey,
				SourceId:        tt.fields.SourceId,
				TeamId:          tt.fields.TeamId,
			}
			gotResp, err := m.VideoCodecList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoCodecList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotResp.Data {
				fmt.Println(v.ID, v.Name)
			}
		})
	}
}

func TestMTeamApi_TeamList(t *testing.T) {
	type fields struct {
		URL             string
		ApiKey          string
		UploadImgApiKey string
		SourceId        string
		TeamId          string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *TeamListResponse
		wantErr  bool
	}{
		{
			fields: fields{
				URL:             "",
				ApiKey:          "",
				UploadImgApiKey: "",
				SourceId:        "",
				TeamId:          "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MTeamApi{
				URL:             tt.fields.URL,
				ApiKey:          tt.fields.ApiKey,
				UploadImgApiKey: tt.fields.UploadImgApiKey,
				SourceId:        tt.fields.SourceId,
				TeamId:          tt.fields.TeamId,
			}
			gotResp, err := m.TeamList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TeamList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotResp.Data {
				fmt.Println(v.ID, v.Name)
			}

		})
	}
}

func TestMTeamApi_StandardList(t *testing.T) {
	type fields struct {
		URL             string
		ApiKey          string
		UploadImgApiKey string
		SourceId        string
		TeamId          string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *StandardListResponse
		wantErr  bool
	}{
		{
			fields: fields{
				URL:             "",
				ApiKey:          "",
				UploadImgApiKey: "",
				SourceId:        "",
				TeamId:          "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MTeamApi{
				URL:             tt.fields.URL,
				ApiKey:          tt.fields.ApiKey,
				UploadImgApiKey: tt.fields.UploadImgApiKey,
				SourceId:        tt.fields.SourceId,
				TeamId:          tt.fields.TeamId,
			}
			gotResp, err := m.StandardList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("StandardList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotResp.Data {
				fmt.Println(v.ID, v.Name)
			}
		})
	}
}

func TestMTeamApi_CategoryList(t *testing.T) {
	type fields struct {
		URL             string
		ApiKey          string
		UploadImgApiKey string
		SourceId        string
		TeamId          string
		CategoryId      string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *CategoryListResponse
		wantErr  bool
	}{
		{
			fields: fields{
				URL:             "",
				ApiKey:          "",
				UploadImgApiKey: "",
				SourceId:        "",
				TeamId:          "",
				CategoryId:      "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MTeamApi{
				URL:             tt.fields.URL,
				ApiKey:          tt.fields.ApiKey,
				UploadImgApiKey: tt.fields.UploadImgApiKey,
				SourceId:        tt.fields.SourceId,
				TeamId:          tt.fields.TeamId,
				CategoryId:      tt.fields.CategoryId,
			}
			_, err := m.CategoryList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestMTeamApi_Send2(t *testing.T) {
	type fields struct {
		URL             string
		ApiKey          string
		UploadImgApiKey string
		SourceId        string
		TeamId          string
		CategoryId      string
		ProxyEnabled    bool
		ProxyAddr       string
	}
	type args struct {
		ctx  context.Context
		path string
		data map[string]string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantBody []byte
		wantErr  bool
	}{
		{
			fields: fields{
				URL:             "",
				ApiKey:          "",
				UploadImgApiKey: "",
				SourceId:        "5",
				TeamId:          "",
				CategoryId:      "402",
				ProxyEnabled:    true,
				ProxyAddr:       "http://127.0.0.1:10810",
			},
			args: args{
				ctx:  context.Background(),
				path: "/api/torrent/createOredit",
				data: map[string]string{
					"category": "402",
					"descr":    "cc",
					"name":     "123",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MTeamApi{
				URL:             tt.fields.URL,
				ApiKey:          tt.fields.ApiKey,
				UploadImgApiKey: tt.fields.UploadImgApiKey,
				SourceId:        tt.fields.SourceId,
				TeamId:          tt.fields.TeamId,
				CategoryId:      tt.fields.CategoryId,
				ProxyEnabled:    tt.fields.ProxyEnabled,
				ProxyAddr:       tt.fields.ProxyAddr,
			}
			gotBody, err := m.SendForm(tt.args.ctx, tt.args.path, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(string(gotBody))
		})
	}
}
