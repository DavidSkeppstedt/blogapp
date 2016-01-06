package api

import (
	"fmt"
	"net/http"
)

func posts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Method)
	fmt.Fprintf(w, "Hello Wolrd")
}
