package datamodels

import (
	"time"

	"github.com/google/uuid"
)

type Objective struct {
	ObjectiveID  uuid.UUID `gorm:"column:objectiveid;primaryKey"`
	CommitmentID string    `gorm:"column:commitmentid"`
	OwnerID      string    `gorm:"column:ownerid"`

	Description     string
	ObjectiveTypeID uuid.UUID `gorm:"column:objectivetypeid"`
	Status          string
	Prioritylevel   string

	Progress        float32
	Creation_date   time.Time
	Lastupdate      time.Time
	Deadline_date   time.Time
	Completion_date *time.Time

	Notes string

	ObjectiveType ObjectiveType `gorm:"foreignKey:ObjectiveTypeID;references:ObjectiveTypeID"`
	KeyResults    []KeyResult   `gorm:"foreignKey:ObjectiveID;references:ObjectiveID"`
}

// TableName overrides the default table name
func (Objective) TableName() string {
	return "objective"
}

// Request
type GetAllObjectivesRequest struct {
	ObjectiveID   string `filter:"objectiveid_eq"`
	CommitmentID  string `filter:"commitmentid_eq"`
	OwnerID       string `filter:"ownerid_eq"`
	Status        string `filter:"status_eq"`
	Prioritylevel string `filter:"prioritylevel_eq"`

	ObjectiveTypeID string `filter:"objectivetype_eq"`

	Description string `filter:"description_ilike"`
	Notes       string `filter:"notes_ilike_ilike"`

	Progress_GTE        string `filter:"progress_gte"`
	Creation_date_GTE   string `filter:"creation_date_gte"`
	Lastupdate_GTE      string `filter:"lastupdate_gte"`
	Deadline_date_GTE   string `filter:"deadline_date_gte"`
	Completion_date_GTE string `filter:"completion_date_gte"`

	Progress_LTE        string `filter:"progress_lte"`
	Creation_date_LTE   string `filter:"creation_date_lte"`
	Lastupdate_LTE      string `filter:"lastupdate_lte"`
	Deadline_date_LTE   string `filter:"deadline_date_lte"`
	Completion_date_LTE string `filter:"completion_date_lte"`
}
