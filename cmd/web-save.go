package cmd

import "github.com/spf13/cobra"

var url string
var useProxy bool

var webSaveCmd = &cobra.Command{
	Use:   "save-web",
	Short: "save web page",
	Long:  "save web page into markdown file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	webSaveCmd.Flags().StringVarP(&url, "url", "", "", "url")
	webSaveCmd.Flags().BoolVarP(&useProxy, "useProxy", "p", false, "url")
}
