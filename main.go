package main

import (
	"log"
	"net/http"
	"os"
)

const DefaultPort = "7893"

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

func echoHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

	writer.Header().Set("Access-Control-Allow-Origin", "*")

	// allow pre-flight headers
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	request.Write(writer)
}

func main() {
	log.Println("starting server, listening on port " + getServerPort())

	http.HandleFunc("/", echoHandler)
	http.ListenAndServe(":"+getServerPort(), nil)
}
