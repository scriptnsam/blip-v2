//go:build !windows

package sect

import (
	"os/exec"
)

func SetSysProcAttr(cmd *exec.Cmd) {
}
