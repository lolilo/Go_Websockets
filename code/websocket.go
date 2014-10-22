// JavaScript to place into browser console to test server code. 

/* 
var sock = new WebSocket("ws://localhost:4000/");
sock.onmessage = function(m) { console.log("Received:", m.data); }
sock.send("Hello!\n")
*/

package code

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
)

func Websocket() {
	http.Handle("/", websocket.Handler(websocket_handler)) // register websocket handler
	http.ListenAndServer("localhost:4000", nil)
}

// 25:21
// bidirectional communication channel
func websocket_handler(c *websocket.Conn) {
	var s string
	fmt.Fscan(c, &s)
	fmt.Println("Received:", s)
	fmt.Fprint(c, "How do you do?")
}

