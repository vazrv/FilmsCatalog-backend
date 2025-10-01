package config

// Файл с настройками.

// Тут хранятся параметры: порт сервера, адрес базы, логин и пароль.
// Эти данные берутся из .env файла или ставятся по умолчанию.

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Структура для хранения всех настроек приложения
type Config struct {
	Env        string
	Port       int
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

// Создание конфигурации с загрузкой значений из переменных окружения
func NewConfig() *Config {
	_ = godotenv.Load() // Загружаем переменные окружения из .env файла

	// Парсим числовые параметры с значениями по умолчанию
	port, _ := strconv.Atoi(getEnv("PORT", "8080"))
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return &Config{
		Env:        getEnv("ENV", "development"),
		Port:       port,
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASS", "root"),
		DBName:     getEnv("DB_NAME", "FilmsCatalog"),
	}
}

// Формирование адреса сервера в формате ":порт"
func (c *Config) ServerAddress() string {
	return ":" + strconv.Itoa(c.Port)
}

// Вспомогательная функция для получения переменных окружения
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
