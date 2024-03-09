// ----------
// Web server
// ----------

// If we write our own web server, we would like to test it as well without manually having to
// stand up a server. The Go standard library also supports this. Below is our simple web server.

package main

import (
	"Gorepo/go_programing/learn_golang/ultimate_go/Testing/example_internal_web_end_point/handlers"
	"log"
	"net/http"
	// Import handler package that has a set of routes that we are gonna work with.
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
