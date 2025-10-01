package middleware

// Файл для логов. Логирует каждый запрос (метод, адрес, время).
// Ловит ошибки и паники, чтобы сервер не падал. Добавляет каждому запросу уникальный ID, чтобы легче искать в логах.

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger структура для логгера
type Logger struct {
	Zap *zap.Logger
}

// NewLogger создает и настраивает новый логгер
func NewLogger(env string) (*Logger, error) {
	var config zap.Config

	if env == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{Zap: logger}, nil
}

// GinLogger middleware для логирования HTTP запросов в gin
func (l *Logger) GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		l.Zap.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("duration", duration),
			zap.String("request_id", c.GetString("request_id")),
		)
	}
}

// GinRecovery middleware для обработки паник в gin
func (l *Logger) GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				l.Zap.Error("Panic recovered",
					zap.Any("error", rec),
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("ip", c.ClientIP()),
					zap.Stack("stack"),
				)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
			}
		}()

		c.Next()
	}
}

// RequestID middleware для добавления ID запроса
func (l *Logger) RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		c.Set("request_id", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()
	}
}

// generateRequestID генерирует уникальный ID запроса
func generateRequestID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		// fallback -- hex timestamp + random bytes
		now := time.Now().UnixNano()
		return hex.EncodeToString([]byte(time.Unix(0, now).Format("20060102150405"))) + "-" + randomHex(8)
	}
	return time.Now().Format("20060102150405") + "-" + hex.EncodeToString(b)
}

func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "rnd"
	}
	return hex.EncodeToString(b)
}

// Sync закрывает логгер (вызывайте при завершении приложения)
func (l *Logger) Sync() {
	_ = l.Zap.Sync()
}

// Sugar возвращает sugared logger для удобного логирования
func (l *Logger) Sugar() *zap.SugaredLogger {
	return l.Zap.Sugar()
}
