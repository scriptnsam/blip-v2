/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/scriptnsam/blip-v2/cmd/me"
	"github.com/scriptnsam/blip-v2/pkg/dependency"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	versionFlag bool
	setupFlag   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "blip",
	Short: "Move easily",
	Long:  `Blip is a command-line interface (CLI) utility designed to streamline the process of migrating applications from one computer to another. This tool is particularly useful for users who are transitioning to a new computer and wish to avoid the hassle of manually reinstalling all their preferred applications. By automating the download and installation process based on a predefined list of applications, Blip simplifies the transition, saving users valuable time and effort.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("v1.2.5")
		} else if setupFlag {
			resp := dependency.SetupChocolatey()
			fmt.Println(resp)
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func AddSubCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blip-v2.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "View current version")
	rootCmd.Flags().BoolVarP(&setupFlag, "setup", "s", false, "Setup the Blip CLI")

	AddSubCommand(me.MeCmd)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".blip-v2" (without extension).
		viper.AddConfigPath(home)
		// viper.SetConfigType("yaml")
		// viper.SetConfigName(".blip-v2")

		// viper.SetConfigFile(".env")

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
