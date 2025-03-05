package data

import (
	"context"
	apimodels "objective-service/api/routes/models"
	datamodels "objective-service/data/models"
	"objective-service/services/postgres"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
	"go.uber.org/fx"
)

func init() {
	core.Provides(fx.Annotate(NewData, fx.As(new(IData))))
}

type IData interface {
	GetPagedObjectives(ctx context.Context, filters *apimodels.GetAllObjectivesRequest) ([]*datamodels.Objective, *core.ApplicationError)
	CountObjectives(ctx context.Context, filters *apimodels.GetAllObjectivesRequest) (int64, *core.ApplicationError)
}

type Data struct {
	PgService *postgres.Service
	// Config *api.Config
}

func NewData(pgService *postgres.Service) *Data {
	return &Data{
		PgService: pgService,
		// Config: apiConfig,
	}
}
