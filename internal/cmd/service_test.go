package cmd

import (
	"context"
	"fmt"
	"pt-auto/internal/utils"
	"testing"
)

func TestService_parseNameByCSharp(t *testing.T) {
	type fields struct {
		TC *utils.TranslateConfig
	}
	type args struct {
		ctx     context.Context
		nameChs string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantInfo *NameInfo
		wantErr  bool
	}{
		{
			name: "",
			fields: fields{
				TC: &utils.TranslateConfig{
					Key:  "",
					Host: "",
				},
			},
			args: args{
				ctx:     context.Background(),
				nameChs: "家有儿女##-##副标题##Home with Kids",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				TC: tt.fields.TC,
			}
			gotInfo, err := service.parseNameByCSharp(tt.args.ctx, tt.args.nameChs)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNameByCSharp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotInfo.NameChs)
			fmt.Println(gotInfo.Year)
			fmt.Println(gotInfo.Subject)
			fmt.Println(gotInfo.NameEng)
			fmt.Println(gotInfo.Flag)
		})
	}
}

func TestService_readDir(t *testing.T) {
	type fields struct {
		TC           *utils.TranslateConfig
		ProxyEnabled bool
		ProxyAddr    string
	}
	type args struct {
		ctx    context.Context
		dir    string
		filter func(name string) bool
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantFilename []string
		wantErr      bool
	}{
		{
			name: "",
			fields: fields{
				TC:           nil,
				ProxyEnabled: false,
				ProxyAddr:    "",
			},
			args: args{
				ctx:    context.Background(),
				dir:    "D:\\Downloads",
				filter: func(name string) bool { return true },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				TC:           tt.fields.TC,
				ProxyEnabled: tt.fields.ProxyEnabled,
				ProxyAddr:    tt.fields.ProxyAddr,
			}
			gotFilename, err := service.readDir(tt.args.ctx, tt.args.dir, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotFilename)
		})
	}
}
