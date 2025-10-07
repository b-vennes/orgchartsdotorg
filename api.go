package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"orgcharts.org/api/pkg/endpoints"
	"orgcharts.org/api/pkg/models"
	"orgcharts.org/api/pkg/state"
)

type FileStatusWithChannel struct {
	request  chan<- models.Empty
	response <-chan map[models.PartialFileRef][]models.FilePart
}

func (f FileStatusWithChannel) GetFileStatuses() map[models.PartialFileRef][]models.FilePart {
	f.request <- models.Empty{}
	return <-f.response
}

type ChartsServiceWithChannels struct {
	request  chan<- models.Empty
	response <-chan []models.Chart
}

func (c ChartsServiceWithChannels) GetCharts() []models.Chart {
	c.request <- models.Empty{}
	return <-c.response
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

	fileState := models.EmptyFileState()
	chartState := models.EmptyChartState()

	uploadsChannel := make(chan models.FilePart)
	startsChannel := make(chan models.PartialFileRef)
	statusRequestsChannel := make(chan models.Empty)
	fileStatusesChannel := make(
		chan map[models.PartialFileRef][]models.FilePart,
	)
	newChartsChannel := make(chan models.UnparsedChart)
	chartsRequestsChannel := make(chan models.Empty)
	chartsResponseChannel := make(chan []models.Chart)

	go state.ManageFileUploads(
		fileState,
		startsChannel,
		uploadsChannel,
		statusRequestsChannel,
		fileStatusesChannel,
		newChartsChannel,
	)

	go state.ManageCharts(
		chartState,
		newChartsChannel,
		chartsRequestsChannel,
		chartsResponseChannel,
	)

	server := http.NewServeMux()
	server.HandleFunc(
		"/initialize-upload",
		endpoints.HandleInitializeFileUpload(startsChannel),
	)
	server.HandleFunc(
		"/upload-part",
		endpoints.HandleUploadPart(uploadsChannel),
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
		endpoints.HandleGetCharts(
			ChartsServiceWithChannels{
				request:  chartsRequestsChannel,
				response: chartsResponseChannel,
			},
		),
	)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)

	if err != nil {
		log.Fatalln("Failed to start server.", err.Error())
	}
}
