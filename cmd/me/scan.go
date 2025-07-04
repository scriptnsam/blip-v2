/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package me

import (
	"fmt"
	"log"

	"github.com/scriptnsam/blip-v2/pkg/tools"
	"github.com/spf13/cobra"
)

var techFlag bool

// loginCmd represents the login command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan your system for available tools",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please wait while we scan your available apps...")
		apps, err := tools.Scanner(techFlag)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Scan completed successfully")
			fmt.Println("Available tools:")
			// for _, app := range apps {
			// 	fmt.Printf("- %s - %s - %s\n", app.Name, app.Command, app.Source)
			// }

			selectedApps, err := tools.PromptMultiSelectApps(apps)
			if err != nil {
				log.Fatalf("Prompt failed %v\n", err)
			}

			fmt.Println("Selected apps:")
			for _, app := range selectedApps {
				fmt.Printf("- %s - %s - %s\n", app.Name, app.Command, app.Source)
			}
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	MeCmd.AddCommand(scanCmd)

	scanCmd.Flags().BoolVar(&techFlag, "all", false, "Include all technical/system packages")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
