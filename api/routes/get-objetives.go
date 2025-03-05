package routes

import (
	"context"
	"maps"
	"net/http"
	apimodels "objective-service/api/routes/models"

	api "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-api"
	"github.com/danielgtaylor/huma/v2"
)

// Init
func init() {
	maps.Copy(GetAllObjectivesResponses, api.DefaultResponses)
}

// Responses
var GetAllObjectivesResponses = map[string]*huma.Response{
	"200": {Ref: "", Description: "OK", Content: nil, Links: nil, Extensions: nil},
}

// Response
type GetAllObjectivesResponse struct {
	Body *apimodels.GetAllObjectivesResponseBody
}

// Request
// type GetAllObjectivesRequest struct {
// 	Body *apimodels.GetAllObjectivesRequest
// }

// Operation
var GetAllObjectivesOperation = huma.Operation{
	OperationID:   "get-all-objectives",
	Method:        http.MethodGet,
	Path:          "/objectives",
	Summary:       "Get all Objectives",
	Description:   "Get all Objectives on the database, filtered by query parameters",
	Tags:          []string{},
	DefaultStatus: http.StatusOK,
	Responses:     GetAllObjectivesResponses,
}

// Function
func (r *Router) GetAllObjectives(ctx context.Context, input *apimodels.GetAllObjectivesRequest) (*GetAllObjectivesResponse, error) {

	responseBody, err := r.Business.GetAllObjectives(ctx, input)
	if err != nil {
		return nil, err
	}

	response := GetAllObjectivesResponse{
		Body: responseBody,
	}

	return &response, nil
}
