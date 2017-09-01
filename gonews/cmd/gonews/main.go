package main

import (
	"log"
	"net/http"

	"github.com/subalgo/gonews/pkg/model"

	"github.com/subalgo/gonews/pkg/app"
)

const (
	//port     = ":8080" //port ดึงมาจาก file config.go แล้ว
	mongoURL = "mongodb://127.0.0.1:27017"
)

//const port = ":8080"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model: %v", err)
	}
	log.Println(port)
	http.ListenAndServe(port, mux)

	/*
		mux := http.NewServeMux()
		app.Mount(mux)
		http.ListenAndServe(port,mux)
	*/
}
