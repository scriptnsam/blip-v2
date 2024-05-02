package dependency

import (
	"fmt"
	"os"
	"os/exec"
)

func UpdateBlip() string {
	fmt.Println("Updating Blip CLI...")
	// execute the update command
	cmd := exec.Command("npm", "update", "-g", "blip-cli")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		// if the update fails, return an error
		fmt.Println(err)
		return "Error updating Blip CLI"
	}
	// if the update succeeds, return nil
	return "Blip CLI updated successfully"
}
