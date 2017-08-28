package main

import (
	"net/http"

	"github.com/subalgo/gonews/pkg/app"
)

//const port = ":8080"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
}
