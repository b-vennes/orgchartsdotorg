package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"orgcharts.org/api/pkg/models"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func makeErrorResponse(error string) ErrorResponse {
	return ErrorResponse{
		Error: error,
	}
}

func setJsonContentType(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func setBadRequestStatus(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

type FileInitializeRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Parts int    `json:"parts"`
}

type FileStatuses interface {
	GetFileStatuses() map[models.PartialFileRef][]models.FilePart
}

type ChartsService interface {
	GetCharts() []models.Chart
}

type FileStatus struct {
	Key   string            `json:"key"`
	Files []models.FilePart `json:"files"`
}

func MakeFileStatuses(state map[models.PartialFileRef][]models.FilePart) []FileStatus {
	result := []FileStatus{}

	for key := range state {
		value := state[key]
		keyString := fmt.Sprint(
			"{ ",
			key.FileID,
			" ",
			key.Name,
			" ",
			key.Parts,
			" }",
		)
		result = append(result, FileStatus{
			Key:   keyString,
			Files: value,
		})
	}

	return result
}

type HTTPHandler = func(http.ResponseWriter, *http.Request)

func HandleInitializeFileUpload(sendInitialize chan<- models.PartialFileRef) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling request to initialize a file upload!")

		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}

		var fileInitialize FileInitializeRequest

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&fileInitialize)

		responseEncoder := json.NewEncoder(w)
		setJsonContentType(w)

		if err != nil {
			errMessage := "Request body did not contain a valid request."
			log.Println("Request body did not contain a valid request.", err.Error())

			responseEncoder.Encode(makeErrorResponse(errMessage))
			setBadRequestStatus(w)

			return
		}

		log.Println("Initializing file upload.", fileInitialize)

		sendInitialize <- models.MakePartialFileRef(
			fileInitialize.ID,
			fileInitialize.Parts,
			fileInitialize.Name,
		)

		responseEncoder.Encode(models.Empty{})
	}
}

func HandleUploadPart(sendFileParts chan<- models.FilePart) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling request to upload org chart!")

		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}

		var filePart models.FilePart

		bodyDecoder := json.NewDecoder(r.Body)

		err := bodyDecoder.Decode(&filePart)
		responseEncoder := json.NewEncoder(w)

		if err != nil {
			errMessage := "Request body did not contain valid file part data."

			log.Println(errMessage, err.Error())

			responseEncoder.Encode(models.Empty{})
			setBadRequestStatus(w)

			return
		}

		setJsonContentType(w)
		responseEncoder.Encode(models.Empty{})

		sendFileParts <- filePart
		log.Println("Handled org chart upload request!")
	}
}

func HandleGetCharts(charts ChartsService) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling request to get all org charts!")

		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		charts := charts.GetCharts()

		err := json.NewEncoder(w).Encode(charts)

		if err != nil {
			errMessage := "Failed to encode charts as JSON."
			log.Println(errMessage, err.Error())
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}
	}
}

func HandleUploadStatuses(fileStatuses FileStatuses) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling request to get file statuses!")

		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		statuses := MakeFileStatuses(
			fileStatuses.GetFileStatuses(),
		)

		setJsonContentType(w)
		encodingErr := json.NewEncoder(w).Encode(statuses)

		if encodingErr != nil {
			log.Println(
				"Error while encoding response for get file statuses:",
				encodingErr.Error(),
			)
			http.Error(
				w,
				"An error occurred while preparing statuses.",
				http.StatusInternalServerError,
			)
			return
		}

		log.Println("Completed handling request to get file statuses!")
	}
}
