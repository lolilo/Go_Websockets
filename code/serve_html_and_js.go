import "html/template"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, listenAddr) // global variable rootTemplate
}

var rootTemplate = template.Must(template.New("root".Parse('
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script>

	websocket = new WebSocket("ws://{{.}}/socket");
	sebsocket.onmessage = onMessage;
	websocket.onclose = onClose;

</html>
'))
