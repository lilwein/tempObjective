package util

// Type ctxKey string
type ctxKey string

// Constant of type ctxKey
const Request ctxKey = "request"
const Method ctxKey = "method"
const Header ctxKey = "header"
const Body ctxKey = "body"
const Url ctxKey = "url"

const Version ctxKey = "version"
const ServerHost ctxKey = "serverHost"
const ServerPort ctxKey = "serverPort"

const DB ctxKey = "db"

const ResponseWriter ctxKey = "responseWriter"
const HttpRequest ctxKey = "httpRequest"

// Datapath
const DataPath = "./data/products.json"

// #########################################
// REGEX:

// Regex for uuid
const UuidRegex = "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"

// QUERY PARAMETERS:
// [query parameter]_[operator]
const (
	ObjectiveID         = "objectiveid"
	CommitmentID        = "commitmentid"
	OwnerID             = "ownerid"
	Description         = "description"
	Objectivetypeid     = "objectivetypeid"
	Status              = "status"
	Prioritylevel       = "prioritylevel"
	Progress            = "progress"
	Progress_GTE        = "progress_gte"
	Progress_LTE        = "progress_lte"
	Creation_date_GTE   = "creation_date_gte"
	Creation_date_LTE   = "creation_date_lte"
	Lastupdate_GTE      = "lastupdate_gte"
	Lastupdate_LTE      = "lastupdate_lte"
	Deadline_date_GTE   = "deadline_date_gte"
	Deadline_date_LTE   = "deadline_date_lte"
	Completion_date_GTE = "completion_date_gte"
	Completion_date_LTE = "completion_date_lte"
	Notes               = "notes"

	Keyresultid   = "keyresultid"
	Kpi           = "kpi"
	KpiTarget_GTE = "kpitarget_gte"
	KpiTarget_LTE = "kpitarget_lte"

	TaskID   = "taskid"
	Taskname = "taskname"

	Objectivetype = "objectivetype"

	QP_PageNumber = "pageNumber"
	QP_PageSize   = "pageSize"
)

const GetByObjective = "getbyobjective"
const GetByOwner = "getbyowner"
const GetByCommitment = "getbycommitment"

const GetByKeyResult = "getbykeyresult"

// #########################################

// Default number of items per page
const PageSize = 10

// DataLayout date
const DataLayout = "2006-01-02"
