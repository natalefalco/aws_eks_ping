package main
import (
	"fmt"
	"net/http"
    "log"
)
func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/v1/ping", PingHandler)
	http.ListenAndServe("0.0.0.0:8080", handler)
}
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Pong`)
    log.Print("said Pong\n")
}
