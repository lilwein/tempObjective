package apimodels

import (
	"objective-service/entities/page"
	"time"

	"github.com/google/uuid"
)

// Response
type GetAllObjectivesResponseBody struct {
	Found          int64               `json:"found"`
	List           []ObjectiveResponse `json:"dataList"`
	PagingMetaData page.Paging         `json:"pagingMetaData"`
}

// Struct for one Objective response to the ListAll request
type ObjectiveResponse struct {
	ObjectiveID  uuid.UUID `json:"objectiveid" gorm:"column:objectiveid;primaryKey"`
	CommitmentID string    `json:"commitmentid" gorm:"column:commitmentid"`
	OwnerID      string    `json:"ownerid" gorm:"column:ownerid"`

	Description   string `json:"description"`
	Status        string `json:"status"`
	Prioritylevel string `json:"prioritylevel"`

	Progress      float32   `json:"progress"`
	Deadline_date time.Time `json:"deadline_date"`

	ObjectiveType string `json:"_objectivetype" gorm:"-"`
	NKeyResults   int64  `json:"_nkeyresults" gorm:"-"`
}

// TableName overrides the default table name
func (ObjectiveResponse) TableName() string {
	return "objective"
}

// Request
type GetAllObjectivesRequest struct {
	ObjectiveID   string `query:"objectiveid"`
	CommitmentID  string `query:"commitmentid"`
	OwnerID       string `query:"ownerid"`
	Status        string `query:"status"`
	Prioritylevel string `query:"prioritylevel"`

	ObjectiveTypeID string `query:"objectivetype"`

	Description string `query:"description"`
	Notes       string `query:"notes"`

	Progress_GTE        string `query:"progress_gte"`
	Creation_date_GTE   string `query:"creation_date_gte"`
	Lastupdate_GTE      string `query:"lastupdate_gte"`
	Deadline_date_GTE   string `query:"deadline_date_gte"`
	Completion_date_GTE string `query:"completion_date_gte"`

	Progress_LTE        string `query:"progress_lte"`
	Creation_date_LTE   string `query:"creation_date_lte"`
	Lastupdate_LTE      string `query:"lastupdate_lte"`
	Deadline_date_LTE   string `query:"deadline_date_lte"`
	Completion_date_LTE string `query:"completion_date_lte"`

	PageNumber int `query:"pagenumber"`
	PageSize   int `query:"pagesize"`
}
