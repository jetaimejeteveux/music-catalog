package spotify

import (
	"github.com/jetaimejeteveux/music-catalog/internal/configs"
	"github.com/jetaimejeteveux/music-catalog/pkg/httpclient"
	"time"
)

type outbond struct {
	cfg         *configs.Config
	client      httpclient.HttpClient
	AccessToken string
	TokenType   string
	ExpiredAt   time.Time
}

func NewSpotifyOutbond(cfg *configs.Config, client httpclient.HttpClient) *outbond {
	return &outbond{
		cfg:    cfg,
		client: client,
	}
}
