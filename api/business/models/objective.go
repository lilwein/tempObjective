package business

import (
	"time"

	"github.com/google/uuid"
)

// Struct for one Objective response to the ListAll request
type Objective struct {
	ObjectiveID  uuid.UUID
	CommitmentID string
	OwnerID      string

	Description   string
	Status        string
	Prioritylevel string

	Progress      float32
	Deadline_date time.Time

	ObjectiveType string
	NKeyResults   int64
}
