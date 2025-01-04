package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/configs"
	membershipHandler "github.com/jetaimejeteveux/music-catalog/internal/handlers/memberships"
	trackHandler "github.com/jetaimejeteveux/music-catalog/internal/handlers/tracks"
	membershipModel "github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	trackActivityModel "github.com/jetaimejeteveux/music-catalog/internal/models/trackactivities"
	membershipRepo "github.com/jetaimejeteveux/music-catalog/internal/repository/memberships"
	"github.com/jetaimejeteveux/music-catalog/internal/repository/spotify"
	trackActivityRepo "github.com/jetaimejeteveux/music-catalog/internal/repository/trackactivities"
	membershipService "github.com/jetaimejeteveux/music-catalog/internal/service/memberships"
	trackService "github.com/jetaimejeteveux/music-catalog/internal/service/tracks"
	"github.com/jetaimejeteveux/music-catalog/pkg/httpclient"
	"github.com/jetaimejeteveux/music-catalog/pkg/internalsql"
	"log"
	"net/http"
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
	db.AutoMigrate(&membershipModel.User{})
	db.AutoMigrate(&trackActivityModel.TrackActivity{})

	r := gin.Default()

	httpClient := httpclient.NewClient(&http.Client{})

	// init repo
	membershipRepo := membershipRepo.NewRepository(db)
	spotifyOutbond := spotify.NewSpotifyOutbond(cfg, httpClient)
	trackactivities := trackActivityRepo.NewRepository(db)

	// init service
	membershipSvc := membershipService.NewService(cfg, membershipRepo)
	trackSvc := trackService.NewService(spotifyOutbond, trackactivities)

	//init handler
	membershipHandler := membershipHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()
	trackHandler := trackHandler.NewHandler(r, trackSvc)
	trackHandler.RegisterRoute()

	r.Run(":" + cfg.Service.Port)
}
