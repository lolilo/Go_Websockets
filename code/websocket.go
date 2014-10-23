// Type JavaScript into browser console to test server code. 

/* 
Construct a websocket connecting to localhost:4000
var sock = new WebSocket("ws://localhost:4000/");

message handler
every message received will be logged in the JS console
sock.onmessage = function(m) { console.log("Received:", m.data); }

send the message hello
sock.send("Hello!\n")
*/

package code

// Go's websocket package piggybacks off the http package.
import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
)

func Websocket() {
	http.Handle("/", websocket.Handler(websocket_handler)) // register websocket handler
	http.ListenAndServe("localhost:4000", nil)
}

// 25:21

// takes in a single websocket connection
// can read and write to websocket connection
// bidirectional communication channel with JS running in the browser
// when the handler returns, the websocket connection shuts down
func websocket_handler(c *websocket.Conn) {
	var s string
	fmt.Fscan(c, &s) // read a \n newline separate string
	fmt.Println("Received:", s) // prints received string to stdoutput
	fmt.Fprint(c, "How do you do?") // writes response to websocket connection -- appears in JS console
	// websocket closes
}

