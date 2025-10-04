// internal/router/router.go
// Определяет маршруты API и подключает middleware.
package router

// Файл с маршрутами (какие URL доступны). Подключает middleware (логи, CORS, ID запроса).
// Есть маршрут /health, чтобы проверить, жив ли сервер. Тут будут добавляться новые маршруты (/films, /actors и т.п.).

// Middleware (промежуточное ПО) - это функции, которые обрабатывают HTTP-запрос ДО того, как он попадет в основной обработчик.
import (
	"github.com/gin-gonic/gin"

	"FilmsCatalog/internal/app/user"
	"FilmsCatalog/internal/config"
	"FilmsCatalog/internal/db"
	"FilmsCatalog/internal/middleware"
)

// EngineWrapper оборачивает *gin.Engine для инъекции зависимостей
type EngineWrapper struct {
	*gin.Engine
	cfg    *config.Config
	logger *middleware.Logger
	db     *db.DB
}

// NewRouter создаёт новый роутер с нужным режимом работы (debug / release)
func NewRouter(cfg *config.Config, logger *middleware.Logger, database *db.DB) *EngineWrapper {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()

	return &EngineWrapper{
		Engine: engine,
		cfg:    cfg,
		logger: logger,
		db:     database,
	}
}

// RegisterRoutes регистрирует все маршруты приложения
func (r *EngineWrapper) RegisterRoutes() {
	// Глобальные middleware
	r.Use(r.logger.GinRecovery())
	r.Use(r.logger.GinLogger())
	r.Use(r.logger.RequestID())
	r.Use(middleware.CORSMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"env":    r.cfg.Env,
			"db":     "connected",
		})
	})

	// Группа /api
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.Use(middleware.JWTAuth()) // Теперь работает!

			userRepo := user.NewUserRepository(r.db.DB) // Исправлено: r.db.DB вместо r.db.SQL
			userService := user.NewUserService(userRepo)
			userHandler := user.NewUserHandler(userService)

			user.RegisterRoutes(auth, userHandler)
		}
	}
}
