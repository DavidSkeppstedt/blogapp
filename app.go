package main

import (
	"github.com/DavidSkeppstedt/blog/api"
	"runtime"
)

//Entrypoint for the api server
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	api.Start()
}
