/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package me

import (
	"github.com/spf13/cobra"
)

// meCmd represents the me command
var MeCmd = &cobra.Command{
	Use:   "me",
	Short: "Me is a command that represent user's profile",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// rootCmd.AddCommand(MeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
