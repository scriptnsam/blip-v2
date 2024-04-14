package tools

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// The `DownloadTool` function downloads a file from a specified URL and saves it to the user's
// Downloads folder with a specific name based on the operating system.
func DownloadTool(url string, toolName string) (string, error) {
	fileURL := strings.TrimSpace(url)

	// Get the user's home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Error:", err)
		return "", err
	}

	// constructy the path to the downloads folder
	downloadFolder := filepath.Join(usr.HomeDir, "Downloads")

	// create a new folder in the downloads folder
	newFolder := filepath.Join(downloadFolder, "blip_tools")
	if err := os.MkdirAll(newFolder, 0755); err != nil {
		log.Fatal("Error creating folder:", err)
		return "", err
	}

	// name of the file to save the downloaded content as
	var fileName string
	switch os := runtime.GOOS; os {
	case "windows":
		fileName = filepath.Join(newFolder, toolName+".exe")
	case "darwin":
		fileName = filepath.Join(newFolder, toolName+".dmg")
	case "linux":
		fileName = filepath.Join(newFolder, toolName+".tar.gz")
	default:
		log.Fatal("Unsupported operating system:", os)
		return "", err

	}

	// Create the file to save the downloaded content
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating file:", err)
		return "", err
	}
	defer file.Close()

	// Perform HTTP request to download the content
	resp, err := http.Get(fileURL)
	if err != nil {
		log.Fatal("Error downloading file:", err)
		return "", err
	}

	defer resp.Body.Close()

	// cherckl if the request was successful
	if resp.StatusCode != http.StatusOK {
		log.Fatal("Error downloading file:", resp.Status)
		return "", err
	}

	// copy the content from the response to the file
	_, err = file.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal("Error copying file:", err)
		return "", err
	}

	return fileName, nil
}

func DownloadAndInstallTool(packageName string) (string, error) {

	// Ensure chocolatey is installed on the system
	if !IsChocolateyInstalled() {
		if err := InstallChocolatey(); err != nil {
			log.Fatal("Error installing chocolatey:", err)
			return "", err
		}
	}

	checkChocolateyVersion()

	// Install desired package using chocolatey
	if err := chocoInstall(packageName); err != nil {
		log.Fatal("Error installing package:", err)
		return "", err
	}
	return "Package " + packageName + " is " + "installed successfully.", nil
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
func checkChocolateyVersion() {
	cmd := exec.Command("choco", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: Chocolatey not found. Check installation.")
	} else {
		fmt.Println("Chocolatey installed successfully.")
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
