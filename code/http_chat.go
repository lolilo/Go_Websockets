// 26:38
// using the http and websocket packages

package code

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

func HTTP_chat() {
	http.HandleFunc("/", rootHandler) // register root handler function with the web root
	http.Handle("/socket", websocket.Handler(socketHandler))
	
	err := http.ListenAndServe(listenAddr, nil)	// web server
	if err != nil {
		log.Fatal(err)
	}

}
