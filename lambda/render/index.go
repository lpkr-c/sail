package render

import (
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"
	"strconv"
	"time"

	"github.com/devinmcgloin/sail/pkg/renderer"
	"github.com/devinmcgloin/sail/pkg/slog"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	slog.SetLevel(slog.DEBUG)
	sketchID := "delaunay/ring" //fmt.Sprintf("%s/%s", ps.ByName("category"), ps.ByName("sketch"))
	seedString := ""            //ps.ByName("seed")
	var seed int64
	if seedString == "" {
		seed = time.Now().Unix()
	} else {
		i, err := strconv.ParseInt(seedString, 0, 64)
		if err != nil {
			seed = hash(seedString)
		} else {
			seed = i
		}
	}
	bytes, err := renderer.Render(sketchID, false, seed)
	if err != nil {
		fmt.Fprintf(w, "An Error Occured: %s\n", err)
		slog.ErrorPrintf("an error occured: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if bytes == nil {
		err := errors.New("bytes is nil")
		fmt.Fprintf(w, "An Error Occured: %s\n", err)
		slog.ErrorPrintf("an error occured: %s\n", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes.Bytes())
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
