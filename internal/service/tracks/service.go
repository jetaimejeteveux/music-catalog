package tracks

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivities"
	"github.com/jetaimejeteveux/music-catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=tracks
type SpotifyOutbond interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type trackActivitiesRepository interface {
	Create(ctx context.Context, model trackactivities.TrackActivity) error
	Update(ctx context.Context, model trackactivities.TrackActivity) error
	Get(ctx context.Context, userId uint, spotifyId string) (*trackactivities.TrackActivity, error)
	GetBulkSpotifyIDs(ctx context.Context, userId uint, spotifyIds []string) (map[string]trackactivities.TrackActivity, error)
}
type service struct {
	spotifyOutbond      SpotifyOutbond
	trackActivitiesRepo trackActivitiesRepository
}

func NewService(spotifyOutbond SpotifyOutbond, repository trackActivitiesRepository) *service {
	return &service{
		spotifyOutbond:      spotifyOutbond,
		trackActivitiesRepo: repository,
	}
}
