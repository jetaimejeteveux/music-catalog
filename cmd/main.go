package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/configs"
	memberships4 "github.com/jetaimejeteveux/music-catalog/internal/handlers/memberships"
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	memberships2 "github.com/jetaimejeteveux/music-catalog/internal/repository/memberships"
	memberships3 "github.com/jetaimejeteveux/music-catalog/internal/service/memberships"
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
	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	// init repo
	membershipRepo := memberships2.NewRepository(db)

	// init service
	membershipSvc := memberships3.NewService(cfg, membershipRepo)

	//init handler
	membershipHandler := memberships4.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(":" + cfg.Service.Port)
}
