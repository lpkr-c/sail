package cmd

import (
	"github.com/devinmcgloin/sail/pkg/library"
	"github.com/spf13/cobra"
)

// sketchesCmd represents the info command
var sketchesCmd = &cobra.Command{
	Use:   "sketches",
	Short: "sketches lists the possible sketches to generate",

	Run: func(cmd *cobra.Command, args []string) {
		regex, _ := cmd.Flags().GetString("regex")
		library.List(regex)
	},
}

func init() {
	rootCmd.AddCommand(sketchesCmd)
	sketchesCmd.PersistentFlags().StringP("regex", "r", ".*", "search for sketches that match this regex")
}
