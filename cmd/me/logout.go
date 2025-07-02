/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package me

import (
	"fmt"

	"github.com/scriptnsam/blip-v2/pkg/authentication"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from your Blip account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := authentication.Logout()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	MeCmd.AddCommand(logoutCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
