package api

import (
	"net/http"
)

func Start() {
	router := NewRouter()
	http.ListenAndServe(":8080", router)
}
