package dependency

import (
	"github.com/scriptnsam/blip-v2/pkg/tools"
)

func SetupChocolatey() string {
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

}
