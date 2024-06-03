package server

import (
	"backend-speaker-clone/internal/cache"
	"backend-speaker-clone/internal/configs"
	"backend-speaker-clone/internal/database"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewServer() {
	configs.LoadEnvFile()

	// --------- datasource init -------------
	database.PostgresConnect()
	// models.MigrateDb()     // update database models
	cache.MemCacheInit()   // cache
	cache.MyRedisConnect() // redis

	// consumers.InitKafkaConsumer() // Kafka

	// kcInstance := utils.InitKeycloak()
	// if kcInstance == nil {
	// 	logger.Warn("Cannot initialize keycloak")
	// }
	// logger.Info("Initialize keycloak successfully")

	// ============ Server start
	// r := controllers.NewRouter()
	log.Println("Server is running ...", "listen", configs.GetModulePort())
	// errRun := r.Run(":" + configs.GetModulePort())
	// if errRun != nil {
	// 	log.Println(errRun)
	// }
}
