// internal/config/config.go
// Управление конфигурацией приложения через переменные окружения (.env).
package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config хранит все настройки приложения
type Config struct {
	Env        string
	HTTPPort   string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

// Глобальный экземпляр конфигурации (синглтон)
var globalConfig *Config

// NewConfig создаёт новый объект конфигурации, загружая значения из .env
func NewConfig() *Config {
	_ = godotenv.Load() // Игнорируем ошибку, если .env нет

	port := getEnv("HTTP_PORT", "8080")
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return &Config{
		Env:        getEnv("ENV", "development"),
		HTTPPort:   port,
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASS", "root"),
		DBName:     getEnv("DB_NAME", "filmscatalog"),
		JWTSecret:  getEnv("JWT_SECRET", "very-secret-key-change-in-production"),
	}
}

// Get возвращает глобальную конфигурацию (синглтон)
func Get() *Config {
	if globalConfig == nil {
		globalConfig = NewConfig()
	}
	return globalConfig
}

// ServerAddress возвращает строку для запуска HTTP-сервера (например, ":8080")
func (c *Config) ServerAddress() string {
	return ":" + c.HTTPPort
}

// Вспомогательная функция для получения переменной окружения
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
