/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package me

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/scriptnsam/blip-v2/pkg/tools"
	"github.com/spf13/cobra"
)

var(
	groupsFlag bool
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "See the list of your available tool",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		if groupsFlag{
			groups,err:=tools.ViewGroups()
			if err!=nil{
				log.Fatal(err)
			}

			// create a slice of strng to store the options for the prompt
			options := make([]string, 0)

			// give serial number for each group
			for i,group:=range groups{
				options = append(options, fmt.Sprintf("%v. %s", i+1, group.Name))
			}

			prompt := promptui.Select{
				Label: "Select a group",
				Items: options,
			}

			_,selectedOption,err:=prompt.Run()
			if err!=nil{
				log.Fatal(err)
			}

			// Check if no option was selected
			if len(selectedOption) == 0 {
				fmt.Println("No option selected.")
				return
			}

			// Find the selected group based on the selected option
			selectedIndex:= selectedOption[0]-'0'-1
			selectedGroup:=groups[selectedIndex]

			// fmt.Printf("You selected: %s\nPlease wait...",selectedGroup.Name)

			t,err:=tools.ViewToolsByGroup(selectedGroup.Name)
			if err!=nil{
				log.Fatal(err)
			}

			fmt.Println(t)

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

				if input=="n"{
					fmt.Println("Exiting...")
					return
				}
				break
			}
			
			// download the tools
			fmt.Printf("Tools download starting...\n\n")
			for _,tool:=range t{
				fmt.Println("Downloading",tool.Name)
				fmt.Println("Please wait...")
				resp,err:=tools.DownloadTool(tool.DownloadLink,tool.Name)
				if err!=nil{
					log.Fatal(err)
				}
				fmt.Printf("Successfully downloaded %v to %s\n\n",tool.Name,resp)
			}

			fmt.Println("Tools download completed.")

		}else{
			resp,err:=tools.ViewTools()
			if err!=nil{
				log.Fatal(err)
			}
			fmt.Println(resp)
		}		
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	viewCmd.Flags().BoolVarP(&groupsFlag,"groups","g",false,"View groups of tools")


	MeCmd.AddCommand(viewCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
