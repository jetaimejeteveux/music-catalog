package tracks

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=tracks
type SpotifyOutbond interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}
type service struct {
	spotifyOutbond SpotifyOutbond
}

func NewService(spotifyOutbond SpotifyOutbond) *service {
	return &service{
		spotifyOutbond: spotifyOutbond,
	}
}
