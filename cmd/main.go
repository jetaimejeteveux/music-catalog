package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/configs"
	"github.com/jetaimejeteveux/music-catalog/pkg/internalsql"
	"log"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("fail to initialize config", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("fail to connect to database %v", err)
	}

	r := gin.Default()
	r.Run(":" + cfg.Service.Port)
}
