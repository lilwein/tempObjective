package postgres

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService(config *Config, lc fx.Lifecycle) *Service {

	service := &Service{}
	psqlInfo := configurePostgres(config)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error

			service.DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{Logger: &ZerologLogger{logger: log.Logger}})
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to connect to Postgres")
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {

			if service.DB != nil {
				log.Info().Msg("Disconnetting Postgres")

				sqlDB, err := service.DB.DB()
				if err != nil {
					log.Fatal().Err(err).Msg("Failed casting from *gorm.DB to *sql.DB")
				}

				sqlDB.Close()
				if err != nil {
					log.Fatal().Err(err).Msg("Failed disconnect Postgres")
				}
				return err
			}
			return nil
		}})

	return service

}

func configurePostgres(cfg *Config) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)

	return psqlInfo
}
