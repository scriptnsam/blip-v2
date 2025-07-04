/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package me

import (
	"fmt"
	"syscall"

	"github.com/scriptnsam/blip-v2/pkg/authentication"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	username string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your Blip account",
	Long:  `You need to be authentcated to peform other action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enter Password:")

		bytePassword, err := term.ReadPassword(int(syscall.Stdin))

		if err != nil {
			fmt.Println(err)
		}

		password := string(bytePassword)

		resp, err := authentication.Login(username, password)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	},
}

func init() {

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username")

	if err := loginCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	// Here you will define your flags and configuration settings.
	MeCmd.AddCommand(loginCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
