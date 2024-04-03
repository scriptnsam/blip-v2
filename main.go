/*
Copyright © 2024 SCRIPTNSAM oluwafemisam40@gmail.com
*/
package main

import (
	"github.com/scriptnsam/blip-v2/cmd"
	"github.com/scriptnsam/blip-v2/pkg/security"
)

func main() {
	// Load environment variable
	security.LoadEnv()
	cmd.Execute()
}
