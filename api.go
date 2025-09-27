package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"orgcharts.org/api/pkg/endpoints"
	"orgcharts.org/api/pkg/models"
)

func manageFileUploads(
	state models.AppState,
	starts <-chan models.PartialFileRef,
	uploads <-chan models.FilePart,
	statusRequests <-chan models.Empty,
	statusResponses chan<- map[models.PartialFileRef][]models.FilePart,
) {
	for {
		select {
		case start := <-starts:
			state.StartUpload(start)
			log.Println(state)
		case upload := <-uploads:
			state.AddPart(upload)
			log.Println(state)
		case <-statusRequests:
			statusResponses <- state.ActiveUploads
		}
	}
}

type FileStatusWithChannel struct {
	request  chan<- models.Empty
	response <-chan map[models.PartialFileRef][]models.FilePart
}

func (f FileStatusWithChannel) GetFileStatuses() map[models.PartialFileRef][]models.FilePart {
	f.request <- models.Empty{}
	return <-f.response
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

	appState := models.EmptyAppState()
	uploadsChannel := make(chan models.FilePart)
	startsChannel := make(chan models.PartialFileRef)
	statusRequestsChannel := make(chan models.Empty)
	fileStatusesChannel := make(
		chan map[models.PartialFileRef][]models.FilePart,
	)

	go manageFileUploads(
		appState,
		startsChannel,
		uploadsChannel,
		statusRequestsChannel,
		fileStatusesChannel,
	)

	server.HandleFunc(
		"/initialize-upload",
		endpoints.HandleInitializeFileUpload(
			endpoints.MakeStartUploadChute(startsChannel),
		),
	)
	server.HandleFunc(
		"/upload-part",
		endpoints.HandleUploadPart(
			endpoints.MakePartsChute(uploadsChannel),
		),
	)
	server.HandleFunc(
		"/upload-status",
		endpoints.HandleUploadStatuses(
			FileStatusWithChannel{
				request:  statusRequestsChannel,
				response: fileStatusesChannel,
			},
		),
	)
	server.HandleFunc(
		"/charts",
		endpoints.HandleGetCharts,
	)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)

	if err != nil {
		log.Fatalln("Failed to start server.", err.Error())
	}
}
