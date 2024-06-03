package models

import (
	"backend-speaker-clone/internal/configs"
	"log"
)

func MigrateDb() {
	// db := database.GetPostgres()
	if configs.GetPostgresIsMigrated() {
		log.Println("migrate db start")
		// ---------- general -----------------
		// db.AutoMigrate(&CustomerInfo{})
		// db.AutoMigrate(&ActionLog{})
		log.Println("Migrate Ended")
	}
}
