package api

import (
	"encoding/json"
	"net/http"
)

type appContext struct {
}

func (c *appContext) helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}
