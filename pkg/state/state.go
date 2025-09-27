package state

import (
  "log"

  "orgcharts.org/api/pkg/files"
  "orgcharts.org/api/pkg/models"
)

func ManageFileUploads(
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
			added := state.AddPart(upload)

			if added == nil {
				log.Println("Uploaded part has no root file.")
				continue
			}

			fileStatus := state.ActiveUploads[*added]

			complete := files.HasAllParts(added.FileID, added.Parts, fileStatus)

			if complete {
				log.Println(
					"Completed upload of file ",
					added.FileID,
					":",
					files.CombineParts(fileStatus),
				)

				updatedCharts := state.Charts

				updatedCharts = append(
					updatedCharts,
					models.MakeChartRef(added.FileID, added.Name),
				)

				state.Charts = updatedCharts
				delete(state.ActiveUploads, *added)
			}

			log.Println(state)
		case <-statusRequests:
			statusResponses <- state.ActiveUploads
		}
	}
}
