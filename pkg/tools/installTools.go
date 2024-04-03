package tools

import (
	"os"
	"os/exec"
	"runtime"
)

func InstallTool(filePath string) (string, error) {
	switch OS := runtime.GOOS; OS {
	case "linux":
		return "sudo apt install -y " + filePath, nil
	case "windows":
		cmd := exec.Command("runas", "/user:Administrator", filePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return "", err
		}
	}

	return "Installation complete", nil
}