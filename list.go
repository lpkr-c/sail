package list

import (
	"fmt"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
