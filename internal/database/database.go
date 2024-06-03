package database

import (
	"backend-speaker-clone/internal/configs"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var dbPostGres *gorm.DB

func PostgresConnect() {
	var err error
	user, pwd := configs.GetPostgresUser(), configs.GetPostgresPassword()
	host, port := configs.GetPostgresHost(), configs.GetPostgresPort()
	dbName := configs.GetPostgresName()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pwd, dbName, port)
	log.Println(dsn)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Enable color
		},
	)

	dbPostGres, err = gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
		},
	), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: configs.GetPostgresSchema() + ".",
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database @ %s:%s", host, port))
	}

	if configs.GetPostgresDebug() {
		dbPostGres = dbPostGres.Debug()
	}
}

func GetPostgres() *gorm.DB {
	return dbPostGres
}
