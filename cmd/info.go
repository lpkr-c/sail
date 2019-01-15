package cmd

import (
	"github.com/devinmcgloin/sail/pkg/library"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "info lists the possible sketches to generate",

	Run: func(cmd *cobra.Command, args []string) {
		regex, _ := cmd.Flags().GetString("regex")
		library.List(regex)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	infoCmd.PersistentFlags().StringP("regex", "r", ".*", "search for sketches that match this regex")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
