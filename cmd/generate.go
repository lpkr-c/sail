package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/devinmcgloin/sail/pkg/renderer"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		level, err := cmd.Flags().GetInt("verbosity")
		if err != nil {
			log.Fatal(err)
		}
		slog.SetLevel(level)
		seed, _ := cmd.Flags().GetInt64("seed")
		if seed <= 0 {
			seed = time.Now().Unix()
		}
		backup, err := cmd.Flags().GetBool("backup")
		if err != nil {
			log.Fatal(err)
		}
		_, err = renderer.Render(args[0], backup, seed)
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
		slog.SetLevel(slog.ERROR)

		iterations, _ := cmd.Flags().GetInt("iterations")
		threads, _ := cmd.Flags().GetInt("threads")
		backup, _ := cmd.Flags().GetBool("backup")
		fmt.Printf("Running for %d iterations with %d threads\n", iterations, threads)
		err := renderer.RenderBulk(args[0], backup, iterations, threads)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {

	viper.SetConfigName(".sail")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(generateBulk)
	generateCmd.PersistentFlags().BoolP("backup", "b", false, "backup to the cloud or not")

	generateCmd.Flags().Int64P("seed", "s", 0, "seed to greate sketch with")
	generateCmd.Flags().IntP("verbosity", "v", slog.INFO, "how verbose to be with logging")

	generateBulk.Flags().IntP("iterations", "i", 300, "number of times to generate the sketch")
	generateBulk.Flags().IntP("threads", "t", 16, "number of threads used to generate the sketch")

	err := viper.ReadInConfig()
	if err != nil {
		slog.ErrorPrintf("Error while reading config: %s\n", err)
	} else {
		slog.DebugPrintf("using config file: %s\n", viper.ConfigFileUsed())
	}
}
