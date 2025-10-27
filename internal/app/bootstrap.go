// backend/internal/app/bootstrap.go
package app

// Файл, который «собирает» приложение.
//
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

// Application представляет собой основной объект приложения, содержащий все ключевые компоненты.
type Application struct {
	Router *router.EngineWrapper
	DB     *db.DB
	Logger *zap.Logger
	Config *config.Config
}

// Bootstrap инициализирует и настраивает всё приложение.
// Он отвечает за:
// 1. Подключение к базе данных.
// 2. Проверку доступности БД (health check).
// 3. Создание и настройку роутера с маршрутами.
// 4. Возврат полного экземпляра приложения.
func Bootstrap(cfg *config.Config, logger *middleware.Logger) (*Application, error) {
	// Шаг 1: Подключение к базе данных
	database, err := db.New(cfg)
	if err != nil {
		logger.Zap.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Шаг 2: Проверка работоспособности базы
	if err := database.HealthCheck(); err != nil {
		logger.Zap.Fatal("Database health check failed", zap.Error(err))
	}
	logger.Zap.Info("Database connected successfully")

	// Шаг 3: Создание роутера и регистрация всех маршрутов
	routerInstance := router.NewRouter(cfg, logger, database)
	routerInstance.RegisterRoutes()

	// Шаг 4: Возврат собранного приложения
	return &Application{
		Router: routerInstance,
		DB:     database,
		Logger: logger.Zap,
		Config: cfg,
	}, nil
}