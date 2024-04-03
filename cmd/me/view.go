/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package me

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/rodaine/table"
	"github.com/scriptnsam/blip-v2/pkg/tools"
	"github.com/spf13/cobra"
)

var (
	groupsFlag bool
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "See the list of your available tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if groupsFlag {
			groups, err := tools.ViewGroups()
			if err != nil {
				log.Fatal(err)
			}

			// create a slice of strng to store the options for the prompt
			options := make([]string, 0)

			// give serial number for each group
			for i, group := range groups {
				options = append(options, fmt.Sprintf("%v. %s", i+1, group.Name))
			}

			prompt := promptui.Select{
				Label: "Select a group",
				Items: options,
			}

			_, selectedOption, err := prompt.Run()
			if err != nil {
				log.Fatal(err)
			}

			// Check if no option was selected
			if len(selectedOption) == 0 {
				fmt.Println("No option selected.")
				return
			}

			// Find the selected group based on the selected option
			selectedIndex := selectedOption[0] - '0' - 1
			selectedGroup := groups[selectedIndex]

			// fmt.Printf("You selected: %s\nPlease wait...",selectedGroup.Name)

			tG, err := tools.ViewToolsByGroup(selectedGroup.Name)
			if err != nil {
				log.Fatal(err)
			}

			// check if there are no tools in the group
			if len(tG) == 0 {
				fmt.Println("No tools found in the group.")
				return
			}

			headerFmt := color.New(color.FgWhite, color.BgCyan).SprintfFunc()

			tbl := table.New("Name", "Group", "Download Link", "Date Created")
			tbl.WithHeaderFormatter(headerFmt)
			tbl.WithFirstColumnFormatter(color.New(color.FgCyan).SprintfFunc())
			tbl.WithPadding(3)
			for _, tool := range tG {
				tbl.AddRow(tool.Name, tool.Group, tool.DownloadLink, tool.DateCreated)
			}

			tbl.Print()

			// run a function to download each tool in the group
			// ask for tools download consent from user as a prompt
			for {
				consentPrompt := promptui.Prompt{
					Label: "Do you want to download these tools [y/n]?",
				}

				input, err := consentPrompt.Run()
				if err != nil {
					log.Fatal(err)
				}

				if input != "y" && input != "n" {
					fmt.Println("invalid input. please enter 'y' or 'n'")
					continue
				}

				if input == "n" {
					fmt.Println("Exiting...")
					return
				}
				break
			}

			// download the tools
			fmt.Printf("Tools download starting...\n\n")
			for _, tool := range tG {
				fmt.Println("Downloading", tool.Name)
				fmt.Println("Please wait...")
				resp, err := tools.DownloadTool(tool.DownloadLink, tool.Name)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Successfully downloaded %v to %s\n\n", tool.Name, resp)

				// Install the tool
				fmt.Println("Installing", tool.Name)
				fmt.Println("Please wait...")
				s, err := tools.InstallTool(resp)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(s)
			}

			fmt.Println("Tools download and installation completed.")

			t, err := tools.ViewToolsByGroup(selectedGroup.Name)
			if err != nil {
				log.Fatal(err)
			}

			// check if there are no tools in the group
			if len(t) == 0 {
				fmt.Println("No tools found in the group.")
				return
			}

			headerFmtN := color.New(color.FgWhite, color.BgCyan).SprintfFunc()

			tblN := table.New("Name", "Group", "Download Link", "Date Created")
			tblN.WithHeaderFormatter(headerFmtN)
			tblN.WithFirstColumnFormatter(color.New(color.FgCyan).SprintfFunc())
			tblN.WithPadding(3)
			for _, tool := range t {
				tblN.AddRow(tool.Name, tool.Group, tool.DownloadLink, tool.DateCreated)
			}

			tblN.Print()

			// run a function to download each tool in the group
			// ask for tools download consent from user as a prompt
			for {
				consentPrompt := promptui.Prompt{
					Label: "Do you want to download these tools [y/n]?",
				}

				input, err := consentPrompt.Run()
				if err != nil {
					log.Fatal(err)
				}

				if input != "y" && input != "n" {
					fmt.Println("invalid input. please enter 'y' or 'n'")
					continue
				}

				if input == "n" {
					fmt.Println("Exiting...")
					return
				}
				break
			}

			// download the tools
			fmt.Printf("Tools download starting...\n\n")
			for _, tool := range t {
				fmt.Println("Downloading", tool.Name)
				fmt.Println("Please wait...")
				resp, err := tools.DownloadTool(tool.DownloadLink, tool.Name)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Successfully downloaded %v to %s\n\n", tool.Name, resp)

				// Install the tool
				fmt.Println("Installing", tool.Name)
				fmt.Println("Please wait...")
				s, err := tools.InstallTool(resp)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(s)
			}

			fmt.Println("Tools download and installation completed.")

		} else {
			resp, err := tools.ViewTools()
			if err != nil {
				log.Fatal(err)
			}

			headerFmt := color.New(color.FgWhite, color.BgCyan).SprintfFunc()

			tbl := table.New("Name", "Group", "Download Link", "Date Created")
			tbl.WithHeaderFormatter(headerFmt)
			tbl.WithFirstColumnFormatter(color.New(color.FgCyan).SprintfFunc())
			tbl.WithPadding(3)
			for _, tool := range resp {
				tbl.AddRow(tool.Name, tool.Group, tool.DownloadLink, tool.DateCreated)
			}

			tbl.Print()

		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	viewCmd.Flags().BoolVarP(&groupsFlag, "groups", "g", false, "View groups of tools")

	MeCmd.AddCommand(viewCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
