package api

import (
	"github.com/DavidSkeppstedt/blog/database"
	"log"
	"net/http"
)

func Start() {
	dbConn, err := database.Open()
	if err != nil {
		panic(err)
	}

	appContext := &appContext{dbConn}
	router := NewRouter(appContext)
	router.ServeFiles("/client/*filepath", http.Dir("./client/static"))
	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}
