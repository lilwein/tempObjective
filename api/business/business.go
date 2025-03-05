package business

import (
	"context"
	apimodels "objective-service/api/routes/models"
	"objective-service/data"

	api "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-api"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
	"go.uber.org/fx"
)

type IBusiness interface {
	GetVersion() string
	GetAllObjectives(ctx context.Context, filters *apimodels.GetAllObjectivesRequest) (*apimodels.GetAllObjectivesResponseBody, *core.ApplicationError)
}

type Logic struct {
	fx.In
	Data   data.IData
	Config *api.Config
}
