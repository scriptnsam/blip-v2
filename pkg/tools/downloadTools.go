package tools

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func DownloadAndInstallTool(packageName string) (string, error) {

	// perform download operation based on the operating system
	switch runtime.GOOS {
	case "windows":
		// Ensure chocolatey is installed on the system
		if !IsChocolateyInstalled() {
			if err := InstallChocolatey(); err != nil {
				log.Fatal("Error installing chocolatey:", err)
				return "", err
			}
		}

		CheckChocolateyVersion()

		// Install desired package using chocolatey
		if err := chocoInstall(packageName); err != nil {
			log.Fatal("Error installing package:", err)
			return "", err
		}
		return "Package " + packageName + " is " + "installed successfully.", nil
	case "linux":
		if err := linuxInstall(packageName); err != nil {
			return "", err
		}
		return "Package " + packageName + " is " + "installed successfully.", nil
	case "darwin":
		if err := brewInstall(packageName); err != nil {
			return "", err
		}
		return "Package " + packageName + " is " + "installed successfully.", nil
	default:
		return "", fmt.Errorf("unsupported operating system")
	}
}

// check if chocolatey is installed on the system
func IsChocolateyInstalled() bool {
	cmd := exec.Command("choco", "list")
	err := cmd.Run()
	return err == nil
}

func InstallChocolatey() error {
	// Install chocolatey on the system
	var cmdName string

	if runtime.GOOS == "windows" {
		// PowerShell command to install Chocolatey with execution policy bypass
		cmdName = "powershell"
	} else {
		return fmt.Errorf("error installing chocolatey: unsupported operating system")
	}

	fmt.Println("Installing chocolatey...")

	// PowerShell command to install Chocolatey
	powerShellCmd := `iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))`

	cmd := exec.Command(cmdName, "-NoProfile", "-InputFormat", "None", "-ExecutionPolicy", "Bypass", "-Command", powerShellCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error installing Chocolatey: %v\n", err)
		fmt.Printf("PowerShell output:\n%s\n", string(output))
		return err
	}

	// Update PATH to include Chocolatey bin directory
	chocoBinDir := `C:\ProgramData\chocolatey\bin`
	if err := updatePath(chocoBinDir); err != nil {
		return fmt.Errorf("error updating PATH: %v", err)
	}

	fmt.Println("Chocolatey installed successfully.")
	return nil

}

func chocoInstall(packageName string) error {
	// Install desired package using chocolatey
	cmd := exec.Command("choco", "install", packageName, "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error installing package %s: %v", packageName, err)
	}
	return nil
}

// Check Chocolatey version after installation
func CheckChocolateyVersion() (string, error) {
	cmd := exec.Command("choco", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: Chocolatey not found. Check installation.")
		return "", fmt.Errorf("chocolatey not found. Check installation")
	} else {
		fmt.Println("Chocolatey installed successfully.")
		return "Chocolatey installed successfully.", nil
	}
}

// updatePath adds the specified directory to the PATH environment variable
func updatePath(newPath string) error {
	pathEnv := os.Getenv("PATH")
	if !strings.Contains(pathEnv, newPath) {
		// Append new path to existing PATH
		newPathEnv := newPath + ";" + pathEnv

		// Update PATH environment variable
		if err := os.Setenv("PATH", newPathEnv); err != nil {
			return err
		}
	}

	return nil
}

func IsBrewInstalled() bool {
	if runtime.GOOS != "darwin" {
		fmt.Println("Brew is only supported on macOS")
	}

	// check if brew executable exist in PATH
	cmd := exec.Command("brew", "--version")
	cmd.Env = os.Environ() // Use current environment variables
	cmd.Stdout = nil       // suppress stdout ouptput
	cmd.Stderr = nil       // suppress stderr ouptput

	err := cmd.Run()
	return err == nil
}

func brewInstall(packageName string) error {
	// Install desired package using brew
	cmd := exec.Command("brew", "install", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error installing package %s: %v", packageName, err)
	}
	return nil
}

func linuxInstall(packageName string) error {
	// Install desired package using apt-get
	cmd := exec.Command("sudo", "apt-get", "install", "-y", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error installing package %s: %v", packageName, err)
	}
	return nil
}
