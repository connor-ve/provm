/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "provm",
	Short: "A version manager for Progress ABL",
	Long: `Provm is a version management tool for Progress ABL that helps you 
switch between multiple installed versions of Progress ABL. 

With Provm, you can set a global default version, switch to specific versions 
for individual projects, and install or list available versions of Progress ABL. 

Examples of usage:

- Set a global version for all projects
- Install a new version of Progress ABL
- List available or installed versions
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.provm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
