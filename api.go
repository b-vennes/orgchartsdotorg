package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func usesJson(w *http.ResponseWriter) {
	(*w).Header().Add("Content-Type", "application/json")
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request to upload org chart!")

	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	usesJson(&w)
	io.WriteString(w, "{}")

	log.Println("Handled org chart upload request!")
}

func main() {
	serverCommand := flag.NewFlagSet("server", flag.ExitOnError)
	parsedPort := serverCommand.Int("port", 5050, "-p 5050")

	parsingErr := serverCommand.Parse(os.Args)

	if parsingErr != nil {
		log.Fatalln("Failed to parse arguments.")
	}

	port := *parsedPort

	log.Printf("Started! Running server on port %d.\n", port)

	server := http.NewServeMux()

	server.HandleFunc("/upload", handleUpload)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)

	if err != nil {
		log.Fatalln("Failed to start server.", err.Error())
	}
}
