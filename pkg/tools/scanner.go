package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/scriptnsam/blip-v2/internal/models"
	"github.com/scriptnsam/blip-v2/pkg/sect"
)

// Valid package names: alphanumeric, hyphens, periods, underscores
var validPackageName = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9.-_]*$`)

// Blacklist meta-packages that aren't user-relevant
var chocoBlacklist = map[string]bool{
	"chocolatey": true, // Exclude the Chocolatey package manager itself
}

func scanAptApps() ([]models.App, error) {
	cmd := exec.Command("apt", "list", "--manual-installed")
	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var apps []models.App
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[installed]") {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				name := strings.Split(parts[0], "/")[0] // e.g., "git/stable"
				apps = append(apps, models.App{
					Name:    name,
					Source:  "apt",
					Command: "apt install " + name,
				})
			}
		}
	}
	return apps, nil

}

// scanChocoApps detects installed Chocolatey packages
func scanChocoApps() ([]models.App, error) {
	cmd := exec.Command("choco", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run choco list: %w", err)
	}

	var apps []models.App
	scanner := bufio.NewScanner(strings.NewReader(string(output)))

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Skip headers, empty lines, and summary lines
		if line == "" || strings.HasPrefix(line, "Chocolatey") || strings.HasSuffix(line, "packages installed.") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) >= 2 {
			name := parts[0]
			version := parts[1]
			apps = append(apps, models.App{
				Name:    fmt.Sprintf("%s@%s", name, version),
				Source:  "choco",
				Command: "choco install " + name,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading choco output: %w", err)
	}

	return apps, nil
}

func scanNpmApps() ([]models.App, error) {
	cmd := exec.Command("npm", "list", "-g", "--depth=0", "--json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run npm list: %w", err)
	}

	var parsed struct {
		Dependencies map[string]struct {
			Version string `json:"version"`
		} `json:"dependencies"`
	}

	if err := json.Unmarshal(output, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse npm output: %w", err)
	}

	var apps []models.App
	for name, dep := range parsed.Dependencies {
		apps = append(apps, models.App{
			Name:    fmt.Sprintf("%s@%s", name, dep.Version),
			Source:  "npm",
			Command: "npm install -g " + name,
		})
	}

	return apps, nil
}

func scanBrewApps() ([]models.App, error) {
	var apps []models.App

	// Use both formula and cask lists
	sources := map[string]string{
		"brew-formula": "brew list --formula",
		"brew-cask":    "brew list --cask",
	}

	for source, cmdLine := range sources {
		parts := strings.Fields(cmdLine)
		cmd := exec.Command(parts[0], parts[1:]...)

		output, err := cmd.Output()
		if err != nil {
			// Continue if one fails (e.g. no casks), only return error if both fail
			continue
		}

		scanner := bufio.NewScanner(strings.NewReader(string(output)))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}

			apps = append(apps, models.App{
				Name:    line,
				Source:  source,
				Command: line,
			})
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return apps, nil
}

func Scanner() ([]models.App, error) {
	switch runtime.GOOS {
	case "windows":
		chocoApps, err1 := scanChocoApps()
		manualApps, err2 := sect.ScanManuallyInstalledApps()

		// Collect errors if any
		if err1 != nil && err2 != nil {
			return nil, fmt.Errorf("choco error: %v; manual error: %v", err1, err2)
		} else if err1 != nil {
			return manualApps, fmt.Errorf("choco error: %v", err1)
		} else if err2 != nil {
			return chocoApps, fmt.Errorf("manual scan error: %v", err2)
		}

		// Merge the two slices
		allApps := append(chocoApps, manualApps...)
		return allApps, nil

	case "linux":
		aptApps, err1 := scanAptApps()
		npmApps, err2 := scanNpmApps()

		if err1 != nil && err2 != nil {
			return nil, fmt.Errorf("apt error: %v; npm error: %v", err1, err2)
		} else if err1 != nil {
			return npmApps, fmt.Errorf("apt error: %v", err1)
		} else if err2 != nil {
			return aptApps, fmt.Errorf("npm error: %v", err2)
		}

		allApps := append(aptApps, npmApps...)
		return allApps, nil
	case "darwin":
		brewApps, err1 := scanBrewApps()
		manualApps, err2 := sect.ScanAllMacApps()

		if err1 != nil && err2 != nil {
			return nil, fmt.Errorf("brew error: %v; manual error: %v", err1, err2)
		} else if err1 != nil {
			return manualApps, fmt.Errorf("brew error: %v", err1)
		} else if err2 != nil {
			return brewApps, fmt.Errorf("manual error: %v", err2)
		}

		allApps := append(brewApps, manualApps...)
		return allApps, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
