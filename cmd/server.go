package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devinmcgloin/sail/handlers/list"
	"github.com/devinmcgloin/sail/handlers/render"
	"github.com/spf13/cobra"
)

// serverCmd represents the info command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server spins up a webserver to generate images on the fly",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt64("port")
		http.HandleFunc("/", list.Handler)
		http.HandleFunc("/render", render.Handler)

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	},
}

func init() {
	serverCmd.Flags().Int64P("port", "p", 8080, "port to bind server responses to")
	rootCmd.AddCommand(serverCmd)
}
