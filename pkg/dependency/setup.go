package dependency

import (
	"runtime"

	"github.com/scriptnsam/blip-v2/pkg/tools"
)

func SetupChocolatey() string {
	switch runtime.GOOS {
	case "windows":
		// check if chocolatey has been installed
		resp, err := tools.CheckChocolateyVersion()
		if err == nil {
			return resp
		}

		// install chocolatey
		if err := tools.InstallChocolatey(); err != nil {
			return err.Error()
		}

		return "Chocolatey installed: Setup completed"
	case "darwin":
		// check if brew has been installed
		r := tools.IsBrewInstalled()
		if r {
			return "Setup completed"
		}
		// retun install brew message
		return `Brew not installed: please run the following command '/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"'`
	case "linux":
		return "Setup completed"
	default:
		return "Unsupported OS"
	}
}
