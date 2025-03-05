package business

import (
	"context"
	apimodels "objective-service/api/routes/models"
	datamodels "objective-service/data/models"
	"objective-service/entities/page"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
)

func (l Logic) GetAllObjectives(ctx context.Context, input *apimodels.GetAllObjectivesRequest) (*apimodels.GetAllObjectivesResponseBody, *core.ApplicationError) {

	var response = &apimodels.GetAllObjectivesResponseBody{}
	filter := &datamodels.GetAllObjectivesFilter{Prioritylevel: input.Prioritylevel}
	//TODO generazione filtro a partire dall'input
	count, errCount := l.Data.CountObjectives(ctx, filter)
	if errCount != nil {
		return nil, errCount
	}

	response.Found = count

	// PAGING
	var offset, limit int
	// Page Size from query param
	// fmt.Println("############################### ", filters.PageSize)
	pageSize, err := page.GetPageSize(input.PageSize)

	// If Page Size is uncorrect, error
	// If Page Size is not defined, return default number of items.
	// If Page Size is 0, return all items.
	// Otherwise, apply paging.
	if err != nil {
		return response, core.BusinessErrorWithError(err)

	} else if pageSize == 0 {
		limit = 0
		offset = 0

	} else {

		// Initialize Paging Metadata
		response.PagingMetaData = page.InitPaging(pageSize, int(response.Found))

		// Selected Page Number from query param
		selectedPage, err := page.GetPageNumber(input.PageNumber)
		if err != nil {
			return response, core.BusinessErrorWithError(err)
		}
		response.PagingMetaData.SetCurrentPage(selectedPage)

		//
		limit = pageSize
		offset = (selectedPage - 1) * pageSize
	}

	items, errItems := l.Data.GetPagedObjectives(ctx, filter, offset, limit)
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
