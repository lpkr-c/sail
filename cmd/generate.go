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

// bulkCmd represents the generate command
var generateBulk = &cobra.Command{
	Use:   "bulk",
	Short: "renders the desired image many times",
	Long: `generate takes the provided sketch and generates an image.
	It uses the given seed, if iterations is provided along with seed,
	seed takes precedence.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		iterations, _ := cmd.Flags().GetInt("iterations")
		threads, _ := cmd.Flags().GetInt("threads")
		fmt.Printf("Running for %d iterations with %d threads\n", iterations, threads)
		err := renderer.RenderBulk(args[0], iterations, threads)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(generateBulk)
	generateCmd.Flags().Int64P("seed", "s", 0, "seed to greate sketch with")
	generateBulk.Flags().IntP("iterations", "i", 300, "number of times to generate the sketch")
	generateBulk.Flags().IntP("threads", "t", 16, "number of threads used to generate the sketch")
}
