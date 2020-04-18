package cmd

import (
	"fmt"

	"github.com/devinmcgloin/sail/pkg/library"
	"github.com/spf13/cobra"
)

// listCmd represents the info command
var listCmd = &cobra.Command{
	Use:     "ls",
	Short:   "lists the possible sketches to generate",
	Aliases: []string{"list", "sketches"},

	Run: func(cmd *cobra.Command, args []string) {
		regex, _ := cmd.Flags().GetString("regex")
		sketches := library.List(regex)
		for _, sketch := range sketches {
			fmt.Printf("%s\n", sketch)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("regex", "r", ".*", "search for sketches that match this regex")
}
