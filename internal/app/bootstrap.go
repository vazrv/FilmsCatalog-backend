package app

// Файл, который «собирает» приложение.

// * Подключает базу.
// * Проверяет, что база доступна.
// * Создаёт роутер (куда какие запросы ходят).
// * Возвращает готовый объект приложения.

import (
	"FilmsCatalog/internal/config"
	"FilmsCatalog/internal/db"
	"FilmsCatalog/internal/middleware"
	"FilmsCatalog/internal/router"

	"go.uber.org/zap"
)

type Application struct {
	Router *router.EngineWrapper
	DB     *db.DB
	Logger *zap.Logger
	Config *config.Config
}

func Bootstrap(cfg *config.Config, logger *middleware.Logger) (*Application, error) {
	database, err := db.New(cfg)
	if err != nil {
		logger.Zap.Fatal("Failed to connect to database", zap.Error(err))
	}

	if err := database.HealthCheck(); err != nil {
		logger.Zap.Fatal("Database health check failed", zap.Error(err))
	}
	logger.Zap.Info("Database connected successfully")

	routerInstance := router.NewRouter(cfg, logger, database)
	routerInstance.RegisterRoutes()

	return &Application{
		Router: routerInstance,
		DB:     database,
		Logger: logger.Zap,
		Config: cfg,
	}, nil
}
