package entities

import "errors"

// Validation errors
var ErrEmptyCommitmentID = errors.New("empty commitment id")
var ErrEmptyOwnerID = errors.New("empty owner id")
var ErrEmptyObjectiveID = errors.New("empty objective id")
var ErrEmptyKeyResultID = errors.New("empty keyresult id")
var ErrEmptyObjectivetypeid = errors.New("empty objectivetype id")

var ErrEmptyDescription = errors.New("empty description")
var ErrEmptyTaskname = errors.New("empty task name")

var ErrInvalidStatus = errors.New("invalid status")
var ErrInvalidPrioritylevel = errors.New("invalid priority level")
var ErrInvalidDeadline_date = errors.New("invalid deadline_date")
