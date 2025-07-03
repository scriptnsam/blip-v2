package sect

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/scriptnsam/blip-v2/internal/models"
)

// 1. GUI Apps (.app bundles)
func scanManualMacApps() ([]models.App, error) {
	dirs := []string{
		"/Applications",
		"/System/Applications",
		filepath.Join(os.Getenv("HOME"), "Applications"),
	}

	seen := make(map[string]bool)
	var apps []models.App

	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() && strings.HasSuffix(entry.Name(), ".app") {
				name := strings.TrimSuffix(entry.Name(), ".app")
				if !seen[name] {
					seen[name] = true
					apps = append(apps, models.App{
						Name:    name,
						Source:  "manual-gui",
						Command: "",
					})
				}
			}
		}
	}
	return apps, nil
}

// 2. .pkg-based installs
func scanPkgUtilApps() ([]models.App, error) {
	cmd := exec.Command("pkgutil", "--pkgs")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var apps []models.App
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		id := scanner.Text()
		apps = append(apps, models.App{
			Name:    id,
			Source:  "manual-pkg",
			Command: "",
		})
	}
	return apps, nil
}

// 3. Mac App Store Apps (requires mas CLI)
func scanMacAppStoreApps() ([]models.App, error) {
	_, err := exec.LookPath("mas")
	if err != nil {
		return nil, nil // mas not installed
	}

	cmd := exec.Command("mas", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var apps []models.App
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 2 {
			apps = append(apps, models.App{
				Name:    strings.TrimSpace(parts[1]),
				Source:  "mac-app-store",
				Command: "",
			})
		}
	}
	return apps, nil
}

// 4. CLI Tools in bin directories
func scanManualCLITools() ([]models.App, error) {
	binDirs := []string{
		"/usr/local/bin",
		"/opt/homebrew/bin",
		filepath.Join(os.Getenv("HOME"), "bin"),
	}
	seen := make(map[string]bool)
	var apps []models.App

	for _, dir := range binDirs {
		erries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, file := range erries {
			if file.Type().IsRegular() || file.Type()&os.ModeSymlink != 0 {
				name := file.Name()
				if !seen[name] {
					seen[name] = true
					apps = append(apps, models.App{
						Name:    name,
						Source:  "manual-cli",
						Command: filepath.Join(dir, name),
					})
				}
			}
		}
	}
	return apps, nil
}

// Combiner for macOS
func ScanAllMacApps() ([]models.App, error) {
	var all []models.App

	sources := []func() ([]models.App, error){
		scanManualMacApps,
		scanPkgUtilApps,
		scanMacAppStoreApps,
		scanManualCLITools,
	}

	for _, f := range sources {
		apps, err := f()
		if err == nil && len(apps) > 0 {
			all = append(all, apps...)
		}
	}

	return all, nil
}
