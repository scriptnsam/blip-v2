/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package me

import (
	"fmt"

	"github.com/scriptnsam/blip-v2/pkg/authentication"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your Blip account",
	Long: `You need to be authentcated to peform other action`,
	Run: func(cmd *cobra.Command, args []string) {
		resp,err := authentication.Login(username, password)

		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(resp)
	},
}

func init() {

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")

	if err:= loginCmd.MarkFlagRequired("username");err!=nil{
		fmt.Println(err)
	}

	if err:= loginCmd.MarkFlagRequired("password");err!=nil{
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
