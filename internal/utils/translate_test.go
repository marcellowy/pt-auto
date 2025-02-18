package utils

import (
	"context"
	"fmt"
	"testing"
)

func TestTranslateText(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *TranslateConfig
		text   string
	}
	tests := []struct {
		name           string
		args           args
		wantTargetText string
		wantErr        bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				config: &TranslateConfig{
					Key:  "", // fill this
					Host: "", // fill this
				},
				text: "中国",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTargetText, err := TranslateText(tt.args.ctx, tt.args.config, tt.args.text, false, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("TranslateText2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotTargetText)
		})
	}
}
