package data

import (
	"context"
	"fmt"
	apimodels "objective-service/api/routes/models"
	datamodels "objective-service/data/models"
	"objective-service/entities/page"
	"objective-service/services"
	"strings"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
	"gorm.io/gorm"
)

// Read all Objectives on db defined in context ctx at key "db"
func (d Data) GetPagedObjectives(ctx context.Context, filters *apimodels.GetAllObjectivesRequest, offset, limit int) ([]*datamodels.Objective, *core.ApplicationError) {

	// DB
	db := d.PgService.DB

	// Response List from DB
	var responseDB = []*datamodels.Objective{}

	// Build query
	var result *gorm.DB = db

	// Preload
	result = result.Preload("ObjectiveType")

	// JOIN
	param := filters.ObjectiveTypeID
	param = strings.ToLower(param)
	if param != "" {
		result = result.Joins("JOIN objectivetype ON objective.objectivetypeid = objectivetype.objectivetypeid")
		result = result.Where("objectivetype.description = ?", param)
	}

	// WHERE: run query
	if err := services.WhereResult(filters, result); err != nil {
		// return response, myerr.NewHttpErr("Error on WhereResult(): ", err, http.StatusBadRequest)
		return nil, err
	}
	result = result.Find(&responseDB)

	// Response
	var responseList = []apimodels.ObjectiveResponse{}

	// Populate Response
	for _, v := range responseDB {
		var keyresults_count int64
		db.Model(&datamodels.KeyResult{}).Where("objectiveid = ?", v.ObjectiveID).Count(&keyresults_count)

		var resp = apimodels.ObjectiveResponse{
			ObjectiveID:  v.ObjectiveID,
			CommitmentID: v.CommitmentID,
			OwnerID:      v.OwnerID,

			Description:   v.Description,
			Status:        v.Status,
			Prioritylevel: v.Prioritylevel,

			Progress:      v.Progress,
			Deadline_date: v.Deadline_date,

			ObjectiveType: v.ObjectiveType.Description,
			NKeyResults:   keyresults_count,
		}
		responseList = append(responseList, resp)
	}

	// PAGING

	// Page Size from query param
	// fmt.Println("############################### ", filters.PageSize)
	pageSize, err := page.GetPageSize(filters.PageSize)

	// If Page Size is uncorrect, error
	// If Page Size is not defined, return default number of items.
	// If Page Size is 0, return all items.
	// Otherwise, apply paging.
	if err != nil {
		return response, core.TechnicalErrorWithError(err)

	} else if pageSize == 0 {
		response.List = responseList

	} else {
		// Initialize Paging Metadata
		response.PagingMetaData = page.InitPaging(pageSize, int(response.Found))

		// Selected Page Number from query param
		selectedPage, err := page.GetPageNumber(filters.PageNumber)
		if err != nil {
			return response, core.TechnicalErrorWithError(err)
		}
		response.PagingMetaData.SetCurrentPage(selectedPage)

		// If Selected Page Number is greater than the available pages,
		// returns an empty page
		if selectedPage > response.PagingMetaData.TotalPages {
			response.List = []apimodels.ObjectiveResponse{}

		} else {
			// Struct Page. Contains list of pages. Each page contains <pageSize> items
			var page = page.Page[apimodels.ObjectiveResponse]{}

			// The response list contains only items of the selected page
			response.List = page.PagingItems(pageSize, selectedPage, responseList, int(response.Found))
		}
	}

	// Errors
	err := result.Error
	if err != nil {
		return nil, core.TechnicalErrorWithCodeAndMessage("QUERY-FAILED", "Error on running query: "+err.Error())
	}

	// Found items
	fmt.Printf("\t\tListAllObjectives() successfull\n\n")
	// fmt.Println(response)

	return responseDB, nil
}

// Read all Objectives on db defined in context ctx at key "db"
func (d Data) CountObjectives(ctx context.Context, filters *apimodels.GetAllObjectivesRequest) (int64, *core.ApplicationError) {

	// Preload
	result := d.PgService.DB.Preload("ObjectiveType")

	// JOIN
	param := filters.ObjectiveTypeID
	param = strings.ToLower(param)
	if param != "" {
		result = result.Joins("JOIN objectivetype ON objective.objectivetypeid = objectivetype.objectivetypeid")
		result = result.Where("objectivetype.description = ?", param)
	}

	// WHERE: run query
	if err := services.WhereResult(filters, result); err != nil {
		// return response, myerr.NewHttpErr("Error on WhereResult(): ", err, http.StatusBadRequest)
		return 0, err
	}
	var count int64
	result = result.Count(&count)

	if result.Error != nil {
		return 0, core.TechnicalErrorWithCodeAndMessage("QUERY-FAILED", "Error on running query: "+result.Error.Error())
	}

	// Found items
	fmt.Printf("\t\tListAllObjectives() successfull\n\n")
	// fmt.Println(response)

	return count, nil
}
