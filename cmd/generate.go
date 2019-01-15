package cmd

import (
	"fmt"
	"time"

	"github.com/devinmcgloin/sail/pkg/renderer"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "renders the desired sketch",
	Long: `generate takes the provided sketch and generates an image.
	It uses the given seed, if iterations is provided along with seed,
	seed takes precedence.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		seed, _ := cmd.Flags().GetInt64("seed")
		if seed <= 0 {
			seed = time.Now().Unix()
		}
		err := renderer.Render(args[0], seed)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")
	generateCmd.PersistentFlags().Int64P("seed", "s", 0, "seed to greate sketch with")

	// Cobra supports local flags which will only run when this command
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
