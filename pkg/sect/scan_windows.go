package sect

import (
	"github.com/scriptnsam/blip-v2/internal/helper"
	"github.com/scriptnsam/blip-v2/internal/models"
	"golang.org/x/sys/windows/registry"
)

func ScanManuallyInstalledApps(showTech bool) ([]models.App, error) {
	var apps []models.App

	keys := []registry.Key{
		registry.LOCAL_MACHINE,
		registry.CURRENT_USER,
	}

	paths := []string{
		`Software\Microsoft\Windows\CurrentVersion\Uninstall`,
		`Software\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
	}

	for _, root := range keys {
		for _, path := range paths {
			k, err := registry.OpenKey(root, path, registry.READ)
			if err != nil {
				continue
			}
			defer k.Close()

			names, err := k.ReadSubKeyNames(-1)
			if err != nil {
				continue
			}

			for _, name := range names {
				subKey, err := registry.OpenKey(k, name, registry.READ)
				if err != nil {
					continue
				}

				displayName, _, err := subKey.GetStringValue("DisplayName")
				if err == nil && displayName != "" {
					if !showTech && helper.ShouldExcludeWindowsApp(displayName) {
						subKey.Close()
						continue
					}

					apps = append(apps, models.App{
						Name:    displayName,
						Source:  "manual",
						Command: "",
					})
				}
				subKey.Close()
			}
		}
	}

	return apps, nil
}
