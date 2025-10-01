package db

// Файл для работы с базой данных.

// Умеет подключаться к PostgreSQL.
// Проверяет соединение.
// Закрывает соединение при завершении работы.
// Настраивает пул соединений (чтобы сервер мог держать много клиентов).

import (
	"fmt"
	"time"

	"FilmsCatalog/internal/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// Обертка вокруг GORM для работы с базой данных
type DB struct {
	*gorm.DB
}

// Создание подключения к базе данных с настройками из конфигурации
func New(cfg *config.Config) (*DB, error) {
	// Формируем строку подключения к PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	// Настраиваем логгер для GORM с фильтрацией по уровню важности
	gormLogger := gormlogger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		gormlogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormlogger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Устанавливаем подключение с оптимизациями производительности
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   gormLogger,
		PrepareStmt:                              true, // кэширование запросов
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Настраиваем пул соединений для эффективной работы с БД
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)           // максимальное количество idle-соединений
	sqlDB.SetMaxOpenConns(100)          // максимальное количество открытых соединений
	sqlDB.SetConnMaxLifetime(time.Hour) // время жизни соединения

	return &DB{DB: gormDB}, nil
}

// Закрытие соединения с базой данных
func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %v", err)
	}
	return sqlDB.Close()
}

// Проверка работоспособности подключения к БД
func (db *DB) HealthCheck() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %v", err)
	}
	return sqlDB.Ping()
}
