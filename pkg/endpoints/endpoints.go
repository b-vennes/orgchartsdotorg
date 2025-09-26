package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"orgcharts.org/api/pkg/models"
)

type ErrorResponse struct {
	Error string
}

type Empty struct{}

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
	ID    string
	Name  string
	Parts int
}

type StartUploadChute struct {
	channel chan<- models.PartialFileRef
}

func MakeStartUploadChute(channel chan<- models.PartialFileRef) StartUploadChute {
	return StartUploadChute{
		channel,
	}
}

func (c *StartUploadChute) send(start models.PartialFileRef) {
	c.channel <- start
}

type PartsChute struct {
	channel chan<- models.FilePart
}

func MakePartsChute(channel chan<- models.FilePart) PartsChute {
	return PartsChute{
		channel: channel,
	}
}

func (p *PartsChute) send(part models.FilePart) {
	p.channel <- part
}

type HTTPHandler = func(http.ResponseWriter, *http.Request)

func HandleInitializeFileUpload(chute StartUploadChute) HTTPHandler {
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

		chute.send(
			models.MakePartialFileRef(
				fileInitialize.ID,
				fileInitialize.Parts,
				fileInitialize.Name,
			),
		)

		responseEncoder.Encode(Empty{})
	}
}

func HandleUploadPart(chute PartsChute) HTTPHandler {
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

			responseEncoder.Encode(Empty{})
			setBadRequestStatus(w)

			return
		}

		setJsonContentType(w)
		responseEncoder.Encode(Empty{})

    chute.send(filePart)
		log.Println("Handled org chart upload request!")
	}
}

func HandleGetCharts(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request to get all org charts!")

	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	charts := []models.ChartRef{
		models.MakeChartRef("1", "First Chart"),
		models.MakeChartRef("2", "Second Chart"),
	}

	err := json.NewEncoder(w).Encode(charts)

	if err != nil {
		errMessage := "Failed to encode charts as JSON."
		log.Println(errMessage, err.Error())
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}
}
