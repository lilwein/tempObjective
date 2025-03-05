package data

import (
	"context"
	"fmt"
	datamodels "objective-service/data/models"
	"objective-service/services"
	"strings"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
	"gorm.io/gorm"
)

// Read all Objectives on db defined in context ctx at key "db"
func (d Data) GetPagedObjectives(ctx context.Context, filters *datamodels.GetAllObjectivesFilter, offset, limit int) ([]*datamodels.Objective, *core.ApplicationError) {

	// DB
	db := d.PgService.DB

	// Response List from DB
	var responseDB = []*datamodels.Objective{}

	// Build query
	var query *gorm.DB = db

	// Preload
	//query = query.Preload("ObjectiveType")

	// JOIN
	param := filters.ObjectiveTypeID
	param = strings.ToLower(param)
	query = query.Joins("ObjectiveType")
	if param != "" {

		query = query.Where("objectivetype.description = ?", param)
	}

	// WHERE: run query
	query, errW := services.WhereResult(query, filters)
	if errW != nil {
		// return response, myerr.NewHttpErr("Error on WhereResult(): ", err, http.StatusBadRequest)
		return nil, errW
	}

	if limit != 0 {
		query = query.Limit(limit).Offset(offset)
	}
	result := query.Find(&responseDB)

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
func (d Data) CountObjectives(ctx context.Context, filters *datamodels.GetAllObjectivesFilter) (int64, *core.ApplicationError) {

	// Preload
	query := d.PgService.DB.Model(&datamodels.Objective{})

	// JOIN
	param := filters.ObjectiveTypeID
	param = strings.ToLower(param)
	if param != "" {
		query = query.Joins("JOIN objectivetype ON objective.objectivetypeid = objectivetype.objectivetypeid")
		query = query.Where("objectivetype.description = ?", param)
	}

	// WHERE: run query
	query, err := services.WhereResult(query, filters)
	if err != nil {
		// return response, myerr.NewHttpErr("Error on WhereResult(): ", err, http.StatusBadRequest)
		return 0, err
	}
	var count int64
	result := query.Count(&count)

	if query.Error != nil {
		return 0, core.TechnicalErrorWithCodeAndMessage("QUERY-FAILED", "Error on running query: "+result.Error.Error())
	}

	// Found items
	fmt.Printf("\t\tListAllObjectives() successfull\n\n")
	// fmt.Println(response)

	return count, nil
}
