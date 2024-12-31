package tracks

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/middleware"
	spotifyModel "github.com/jetaimejeteveux/music-catalog/internal/models/spotify"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotifyModel.SearchResponse, error)
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	router := h.Group("tracks")
	router.Use(middleware.AuthMiddleware())
	router.GET("/search", h.Search)
}