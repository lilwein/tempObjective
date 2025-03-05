package datamodels

import (
	"github.com/google/uuid"
)

// type Objective map[string]interface{}

type ObjectiveType struct {
	ObjectiveTypeID uuid.UUID `json:"objectivetypeid" gorm:"column:objectivetypeid;primaryKey"`
	Description     string    `json:"description"`
}

// TableName overrides the default table name
func (ObjectiveType) TableName() string {
	return "objectivetype"
}
