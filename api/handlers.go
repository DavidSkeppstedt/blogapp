package api

import (
	"encoding/json"
	"github.com/DavidSkeppstedt/blog/database"
	"github.com/DavidSkeppstedt/blog/models"
	_ "github.com/gorilla/context"
	_ "github.com/julienschmidt/httprouter"
	"net/http"
)

type appContext struct {
	db *database.DbConn
}

func (c *appContext) postsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var count string
	if len(params) > 0 && params.Get("count") != "" {
		count = params.Get("count")
	} else {
		count = "100"
	}
	responseData, err := c.db.ListPosts(count)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&models.PostsCollection{responseData})
}
