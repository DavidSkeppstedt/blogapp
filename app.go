package main

import (
	"github.com/DavidSkeppstedt/blog/database"
)

//Entrypoint for the api server
func main() {
	database.Open()
}
