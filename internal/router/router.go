package router

// Файл с маршрутами (какие URL доступны). Подключает middleware (логи, CORS, ID запроса).
// Есть маршрут /health, чтобы проверить, жив ли сервер. Тут будут добавляться новые маршруты (/films, /actors и т.п.).

// Middleware (промежуточное ПО) - это функции, которые обрабатывают HTTP-запрос ДО того, как он попадет в основной обработчик.
import (
	"github.com/gin-gonic/gin"

	"FilmsCatalog/internal/config"
	"FilmsCatalog/internal/db"
	"FilmsCatalog/internal/middleware"
)

// Обертка вокруг роутера Gin для добавления кастомных зависимостей
type EngineWrapper struct {
	*gin.Engine
	cfg    *config.Config
	logger *middleware.Logger
	db     *db.DB
}

// Создание нового роутера с настройкой режима работы
func NewRouter(cfg *config.Config, logger *middleware.Logger, database *db.DB) *EngineWrapper {
	// Устанавливаем режим работы Gin в зависимости от окружения
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

// Регистрация всех middleware и маршрутов приложения
func (r *EngineWrapper) RegisterRoutes() {
	// Подключаем цепочку middleware для обработки запросов
	r.Use(r.logger.GinRecovery())
	r.Use(r.logger.GinLogger())
	r.Use(r.logger.RequestID())
	r.Use(middleware.CORSMiddleware())

	// Health check endpoint для мониторинга состояния приложения
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"env":    r.cfg.Env,
			"db":     "connected",
		})
	})
}
