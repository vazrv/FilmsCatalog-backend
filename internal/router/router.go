// internal/router/router.go
package router

// подключение всех ендпоинтов и создается маршрутизация. ендпоинт адрес на котором лежат все данные,
// по которому бэк взаимодействует с фронтом

import (
	"github.com/gin-gonic/gin"

	"FilmsCatalog/internal/app/actor"
	"FilmsCatalog/internal/app/auth"
	"FilmsCatalog/internal/app/film"
	"FilmsCatalog/internal/app/genre"
	"FilmsCatalog/internal/app/user"
	"FilmsCatalog/internal/config"
	"FilmsCatalog/internal/db"
	"FilmsCatalog/internal/middleware"
)

type EngineWrapper struct {
	*gin.Engine
	cfg    *config.Config
	logger *middleware.Logger
	db     *db.DB
}

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

func (r *EngineWrapper) RegisterRoutes() {
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

	api := r.Group("/api")
	{
		// === Auth: Public Routes ===
		authPublic := api.Group("/auth")
		{
			authRepo := auth.NewUserRepository(r.db.DB)
			authService := auth.NewAuthService(authRepo)
			authHandler := auth.NewAuthHandler(authService)
			auth.RegisterRoutes(authPublic, authHandler)
		}

		// === Auth: Protected Routes (Profile) ===
		authProtected := api.Group("/auth")
		{
			authProtected.Use(middleware.JWTAuth())

			userRepo := user.NewUserRepository(r.db.DB)
			userService := user.NewUserService(userRepo)
			userHandler := user.NewUserHandler(userService)
			user.RegisterRoutes(authProtected, userHandler)
		}

		// === Films API ===
		films := api.Group("/films")
		{
			repo := film.NewRepository(r.db.DB)
			service := film.NewService(repo)
			handler := film.NewHandler(service)
			film.RegisterRoutes(films, handler)
		}

		// === Genres API ===
		genres := api.Group("/genres")
		{
			repo := genre.NewRepository(r.db.DB)
			service := genre.NewService(repo)
			handler := genre.NewHandler(service)
			genre.RegisterRoutes(genres, handler)
		}

		// === Actors API ===
		actors := api.Group("/actors")
		{
			repo := actor.NewRepository(r.db.DB)
			service := actor.NewService(repo)
			handler := actor.NewHandler(service)
			actor.RegisterRoutes(actors, handler)
		}
	}
}
