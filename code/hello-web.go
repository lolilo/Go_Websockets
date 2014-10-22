package code

import (
	"fmt"
	"log"
	"net/http"
)

const listenAddr = "localhost:4000"

func Hello() {

	// register "handler" function with web root "/" 
	// create default route in http package's default router
	http.HandleFunc("/", handler) 

	// listen on specified port "listenAddr" and serve http requests there
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web!")
}
