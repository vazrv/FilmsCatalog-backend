// точка входа в проект, главный файл, запускает сервер. Подгружает настройки (порт, база).
// Включает логгер (чтобы писать сообщения в консоль). Собирает всё приложение через bootstrap.

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"FilmsCatalog/internal/app"
	"FilmsCatalog/internal/config"
	"FilmsCatalog/internal/middleware"

	"go.uber.org/zap"
)

func main() {
	// Инициализация конфигурации и логгера
	cfg := config.NewConfig()
	logger, err := middleware.NewLogger(cfg.Env)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Создание основного приложения
	appInstance, err := app.Bootstrap(cfg, logger)
	if err != nil {
		logger.Zap.Fatal("Application bootstrap failed", zap.Error(err))
	}

	// Настройка HTTP-сервера с таймаутами
	server := &http.Server{
		Addr:         cfg.ServerAddress(),
		Handler:      appInstance.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Запуск сервера в отдельной горутине
	go func() {
		logger.Zap.Info("Server is starting", zap.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Zap.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	// Ожидание сигналов завершения от операционной системы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown - аккуратное завершение работы
	logger.Zap.Info("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Остановка сервера с таймаутом
	if err := server.Shutdown(ctx); err != nil {
		logger.Zap.Fatal("Server forced to shutdown", zap.Error(err))
	}

	// Закрытие подключения к базе данных
	if appInstance.DB != nil {
		if err := appInstance.DB.Close(); err != nil {
			logger.Zap.Error("Failed to close database", zap.Error(err))
		}
	}

	logger.Zap.Info("Server stopped")
}
