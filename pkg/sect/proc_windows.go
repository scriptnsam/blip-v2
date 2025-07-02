package sect

import (
	"os/exec"
	"syscall"
)

func SetSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
