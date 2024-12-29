package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=memberships
type service interface {
	Signup(request memberships.SignupRequest) error
	Login(req memberships.LoginRequest) (string, error)
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
	router := h.Group("memberships")
	router.POST("/signup", h.Signup)
	router.POST("/login", h.Login)
}
