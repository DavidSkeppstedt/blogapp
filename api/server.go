package api

import (
	"github.com/DavidSkeppstedt/blog/database"
	"net/http"
)

func Start() {
	dbConn, err := database.Open()
	if err != nil {
		panic(err)
	}

	appContext := &appContext{dbConn}
	router := NewRouter(appContext)
	http.ListenAndServe(":8080", router)
}
