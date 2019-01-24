package cmd

import (
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
	"time"

	"github.com/devinmcgloin/sail/pkg/renderer"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
)

// serverCmd represents the info command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server spins up a webserver to generate images on the fly",
	Run: func(cmd *cobra.Command, args []string) {
		slog.SetLevel(slog.ERROR)
		router := httprouter.New()
		router.GET("/", index)
		router.GET("/render/:category/:sketch", render)
		router.GET("/render/:category/:sketch/:seed", render)
		fs := http.Dir("sketches")
		router.ServeFiles("/view/*filepath", fs)

		log.Fatal(http.ListenAndServe(":8080", router))
	},
}

func render(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sketchID := fmt.Sprintf("%s/%s", ps.ByName("category"), ps.ByName("sketch"))
	seedString := ps.ByName("seed")
	var seed int64
	if seedString == "" {
		seed = time.Now().Unix()
	} else {
		seed = hash(seedString)
	}
	log.Printf("Rendering %s with seed %d\n", sketchID, seed)
	bytes, err := renderer.Render(sketchID, true, seed)
	if err != nil {
		fmt.Fprintf(w, "An Error Occured: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes.Bytes())
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func hash(s string) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	sum := h.Sum64()
	if int64(sum) <= 0 {
		return int64(sum) * -1
	}
	return int64(sum)
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
