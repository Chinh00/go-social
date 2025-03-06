package database

import (
	"fmt"
	"github.com/google/wire"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var PgContextSet = wire.NewSet(NewPgContext)

type pgContext struct {
	db     *gorm.DB
	config *PostgresConfig
}

func NewPgContext(config *PostgresConfig) DataContext {
	return &pgContext{
		config: config,
	}
}

func (pgContext *pgContext) Migrations(models ...interface{}) error {
	err := pgContext.GetDB().AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}

func (pgContext *pgContext) GetDB() *gorm.DB {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", pgContext.config.Host, pgContext.config.Username, pgContext.config.Password, pgContext.config.Database, pgContext.config.Port)
	db, err := gorm.Open(pg.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}
	pgContext.db = db
	return db
}

func (pgContext *pgContext) Configure(opts ...Options) {
	for _, opt := range opts {
		opt(pgContext)
	}
}
