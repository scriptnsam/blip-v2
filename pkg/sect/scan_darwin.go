package sect

import (
	"fmt"

	"github.com/scriptnsam/blip-v2/internal/models"
)

func ScanManuallyInstalledApps(_ bool) ([]models.App, error) {
	return nil, fmt.Errorf("not supported")
}
