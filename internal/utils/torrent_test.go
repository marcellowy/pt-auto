package utils

import (
	"context"
	"testing"
)

func TestCreateTorrent(t *testing.T) {
	type args struct {
		ctx  context.Context
		root string
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test directory",
			args: args{
				ctx:  context.Background(),
				root: "../utils",
				file: "test1.torrent",
			},
		},
		{
			name: "test file",
			args: args{
				ctx:  context.Background(),
				root: "./torrent_test.go",
				file: "test2.torrent",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTorrent(tt.args.ctx, tt.args.root, tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("CreateTorrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
