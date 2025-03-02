package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

var subCommands = []*cobra.Command{
	wordCmd,
	timeCmd,
	webSaveCmd,
	sqlCmd,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	for _, subCommand := range subCommands {
		rootCmd.AddCommand(subCommand)
	}
}
