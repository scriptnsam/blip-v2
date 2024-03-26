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

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "See the list of your available tool",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		resp,err:=tools.ViewTools()
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println(resp)
	},
}

func init() {
	MeCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
