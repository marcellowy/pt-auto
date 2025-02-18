package utils

import (
	"context"
	"testing"
)

func TestExec(t *testing.T) {
	type args struct {
		ctx   context.Context
		shell string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "test",
			args: args{
				ctx:   context.Background(),
				shell: "dir",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Exec(tt.args.ctx, tt.args.shell)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
