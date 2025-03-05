package business

import (
	"context"
	apimodels "objective-service/api/routes/models"
	datamodels "objective-service/data/models"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
)

func (l Logic) GetAllObjectives(ctx context.Context, input *apimodels.GetAllObjectivesRequest) (*apimodels.GetAllObjectivesResponseBody, *core.ApplicationError) {

	var response = &apimodels.GetAllObjectivesResponseBody{}
	filter := &datamodels.GetAllObjectivesFilter{}
	//TODO generazione filtro a partire dall'input
	count, errCount := l.Data.CountObjectives(ctx, filter)
	if errCount != nil {
		return nil, errCount
	}

	response.Found = count

	items, errItems := l.Data.GetPagedObjectives(ctx, filter, input.PageNumber*input.PageSize, input.PageSize)
	if errItems != nil {
		return nil, errItems
	}

	var list []apimodels.ObjectiveResponse
	for _, v := range items {
		var item = apimodels.ObjectiveResponse{
			ObjectiveID:   v.ObjectiveID,
			CommitmentID:  v.CommitmentID,
			OwnerID:       v.OwnerID,
			Description:   v.Description,
			Status:        v.Status,
			Prioritylevel: v.Prioritylevel,
			Progress:      v.Progress,
			Deadline_date: v.Deadline_date,
			ObjectiveType: v.ObjectiveType.Description,
			NKeyResults:   int64(len(v.KeyResults)),
		}
		list = append(list, item)
	}
	response.List = list

	return response, nil
}
