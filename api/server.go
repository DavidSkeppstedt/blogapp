package api

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/posts", posts)
	http.ListenAndServe(":8080", nil)
}
