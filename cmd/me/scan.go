/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package me

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/scriptnsam/blip-v2/pkg/tools"
	"github.com/spf13/cobra"
)

var (
	techFlag    bool
	shareFlag   bool
	noShareFlag bool
)

// loginCmd represents the login command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan your system for available tools",
	Long:  `This scan collects a list of installed apps. No personal data or device identifiers are stored. Data may be anonymously used to improve community-driven app catalogs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📦 This tool can scan your installed applications.")
		fmt.Println("🔒 No personal data, device IDs, or files will be collected.")
		fmt.Println("📤 You can choose to anonymously share app names to improve the open community catalog.")

		prompt := &promptui.Prompt{
			Label:     "Would you like to continue with the scan?",
			IsConfirm: true,
		}

		confirm, err := prompt.Run()
		if err != nil || strings.ToLower(confirm) != "y" {
			fmt.Println("❌ Scan cancelled.")
			return
		}

		fmt.Println("Please wait while we scan your available apps...")
		apps, err := tools.Scanner(techFlag)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("✅ Scan completed successfully\n")
		fmt.Println("Available tools:")

		selectedApps, err := tools.PromptMultiSelectApps(apps)
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		var share bool

		if shareFlag && noShareFlag {
			// conflicting flags
			log.Println("⚠️ Both --share and --no-share were set. Defaulting to no share.")
			share = false
		} else if shareFlag {
			share = true
			fmt.Println("✅ You opted in via --share flag to contribute anonymously.")
		} else if noShareFlag {
			share = false
			fmt.Println("🔒 You opted out via --no-share flag. Data will remain private.")
		} else {
			// No flag set, prompt user
			sharePrompt := promptui.Select{
				Label: "Do you want to contribute your app list anonymously to the global app index?",
				Items: []string{"Yes, share anonymously", "No, keep it private"},
			}

			_, result, err := sharePrompt.Run()
			if err != nil {
				log.Fatalf("Prompt failed %v\n", err)
			}
			share = strings.HasPrefix(result, "Yes")
		}

		if share {
			fmt.Println("✅ Your app list will be shared anonymously. Thank you for contributing!")
			// Save to global app catalog
		} else {
			fmt.Println("🔒 Your data will remain private.")
		}

		fmt.Println("Selected apps:")
		for _, app := range selectedApps {
			fmt.Printf("- %s (%s)\n", app.Name, app.Source)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	MeCmd.AddCommand(scanCmd)

	scanCmd.Flags().BoolVar(&techFlag, "all", false, "Include all technical/system packages")
	scanCmd.Flags().BoolVar(&shareFlag, "share", false, "Silently agrees to share anonymously")
	scanCmd.Flags().BoolVar(&noShareFlag, "no-share", false, "Overrides prompts")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
