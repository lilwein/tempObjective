package routes

import (
	"context"
	"maps"
	"net/http"

	api "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-api"
	"github.com/danielgtaylor/huma/v2"
)

// Init
func init() {
	maps.Copy(GetVersionResponses, api.DefaultResponses)
}

// Responses
var GetVersionResponses = map[string]*huma.Response{
	"200": {Ref: "", Description: "OK", Content: nil, Links: nil, Extensions: nil},
}

// Operation
var GetVersionOperation = huma.Operation{
	OperationID:   "get-version",
	Method:        http.MethodGet,
	Path:          "/version",
	Summary:       "Dettaglio Versione",
	Description:   "Visualizza versione",
	Tags:          []string{},
	DefaultStatus: http.StatusOK,
	Responses:     GetVersionResponses,
}

// Request
type GetVersionRequest struct {
}

// Response
type GetVersionResponse struct {
	Body *GetVersionResponseBody
}

// Response Body
type GetVersionResponseBody struct {
	Version string
}

// Function
func (r *Router) GetVersion(ctx context.Context, input *GetVersionRequest) (*GetVersionResponse, error) {

	version := r.Business.GetVersion()
	body := &GetVersionResponseBody{
		Version: version,
	}

	resp := GetVersionResponse{Body: body}

	return &resp, nil
}
