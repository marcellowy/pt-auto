package utils

import (
	"context"
	"fmt"
	"github.com/marcellowy/go-common/gogf/vlog"
	"github.com/marcellowy/go-common/os/vstrings"
	"os/exec"
	"syscall"
)

func Exec(ctx context.Context, shell string) (output string, err error) {
	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c %s`, shell), HideWindow: true}
	var outputTmp []byte
	outputTmp, err = cmd.CombinedOutput()
	output = vstrings.GBKToUTF8(string(outputTmp))
	if err != nil {
		vlog.Info(ctx, output)
	}
	return output, err
}
