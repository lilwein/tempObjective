package datamodels

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	TaskID      uuid.UUID `json:"taskid" gorm:"column:taskid;primaryKey"`
	KeyResultID uuid.UUID `json:"keyresultid" gorm:"column:keyresultid"`

	Taskname    string `json:"taskname"`
	Description string `json:"description"`
	Status      string `json:"status"`

	Lastupdate time.Time `json:"lastupdate"`

	Notes string `json:"notes"`

	KeyResult KeyResult `json:"_keyresultid" gorm:"foreignKey:KeyResultID;references:KeyResultID"`
}

// TableName overrides the default table name
func (Task) TableName() string {
	return "task"
}
