package state

import (
	"log"

	"orgcharts.org/api/pkg/files"
	"orgcharts.org/api/pkg/models"
)

func ManageCharts(
	state models.ChartState,
	newCharts <-chan models.UnparsedChart,
	chartRequests <-chan models.Empty,
	chartResponses chan<- []models.Chart,
) {
	for {
		select {
		case newChart := <-newCharts:
			parsedChart := models.MakeChart(
				models.MakeChartRef(
					newChart.ID,
					newChart.Name,
				),
				[]models.Link{},
				[]models.Person{},
			)
			charts := state.Charts
			charts = append(charts, parsedChart)
			state.Charts = charts
		case <-chartRequests:
			chartResponses <- state.Charts
		}
	}
}

func ManageFileUploads(
	state models.FileState,
	starts <-chan models.PartialFileRef,
	uploads <-chan models.FilePart,
	statusRequests <-chan models.Empty,
	statusResponses chan<- map[models.PartialFileRef][]models.FilePart,
	newCharts chan<- models.UnparsedChart,
) {
	for {
		select {
		case start := <-starts:
			state.StartUpload(start)
			log.Println(state)
		case upload := <-uploads:
			added := state.AddPart(upload)

			if added == nil {
				log.Println("Uploaded part has no root file.")
				continue
			}

			fileStatus := state.ActiveUploads[*added]

			complete := files.HasAllParts(added.FileID, added.Parts, fileStatus)
			content := files.CombineParts(fileStatus)

			if complete {
				log.Println(
					"Completed upload of file ",
					added.FileID,
					":",
					content,
				)

				newChart := models.MakeUnparsedChart(added.FileID, added.Name, content)

				newCharts <- newChart

				delete(state.ActiveUploads, *added)
			}

			log.Println(state)
		case <-statusRequests:
			statusResponses <- state.ActiveUploads
		}
	}
}
