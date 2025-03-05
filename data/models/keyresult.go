package datamodels

import (
	"time"

	"github.com/google/uuid"
)

type KeyResult struct {
	KeyResultID uuid.UUID `json:"keyresultid" gorm:"column:keyresultid;primaryKey"`
	ObjectiveID uuid.UUID `json:"objectiveid" gorm:"column:objectiveid"`

	Description   string `json:"description"`
	Status        string `json:"status"`
	Prioritylevel string `json:"prioritylevel"`

	Progress   float32   `json:"progress"`
	Lastupdate time.Time `json:"lastupdate"`

	KPI       string  `json:"kpi" gorm:"column:kpi"`
	KpiTarget float32 `json:"kpitarget" gorm:"column:kpitarget"`

	Notes string `json:"notes"`

	Objective Objective `json:"_objectiveid" gorm:"foreignKey:ObjectiveID;references:ObjectiveID"`
	Tasks     []Task    `json:"_tasks" gorm:"foreignKey:KeyResultID;references:KeyResultID"`
}

// TableName overrides the default table name
func (KeyResult) TableName() string {
	return "key_result"
}
