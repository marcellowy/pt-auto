package utils

func Exec(ctx context.Context, shell string) (output string, err error) {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`-c %s`, shell), HideWindow: true}
	var outputTmp []byte
	outputTmp, err = cmd.CombinedOutput()
	output = string(outputTmp)
	if err != nil {
		vlog.Info(ctx, output)
	}
	return output, err
}
