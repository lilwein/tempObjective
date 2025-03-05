package data

import (
	"context"
	datamodels "objective-service/data/models"
	"objective-service/services/postgres"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
	"go.uber.org/fx"
)

func init() {
	core.Provides(fx.Annotate(NewData, fx.As(new(IData))))
}

type IData interface {
	GetPagedObjectives(ctx context.Context, filters *datamodels.GetAllObjectivesFilter, offset, limit int) ([]*datamodels.Objective, *core.ApplicationError)
	CountObjectives(ctx context.Context, filters *datamodels.GetAllObjectivesFilter) (int64, *core.ApplicationError)
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
