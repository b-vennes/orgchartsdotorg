package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"orgcharts.org/api/pkg/models"
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

func handleGetCharts(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request to get all org charts!")

	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	usesJson(&w)

	charts := []models.ChartRef {
		models.MakeChartRef("1", "First Chart"),
		models.MakeChartRef("2", "Second Chart"),
	}

	usesJson(&w)
	err := json.NewEncoder(w).Encode(charts)

	if err != nil {
		log.Println("Failed to encode charts as JSON.", err.Error())
		http.Error(w, "Failed to encode response as JSON.", http.StatusInternalServerError)
		return
	}
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
	server.HandleFunc("/charts", handleGetCharts)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)

	if err != nil {
		log.Fatalln("Failed to start server.", err.Error())
	}
}
