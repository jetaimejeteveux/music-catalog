package trackactivities

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivities"
)

func (r *repository) GetBulkSpotifyIDs(ctx context.Context, userId uint, spotifyIds []string) (map[string]trackactivities.TrackActivity, error) {
	var models []trackactivities.TrackActivity
	res := r.db.Where("user_id = ?", userId).Where("spotify_id IN (?)", spotifyIds).Find(&models)
	if err := res.Error; err != nil {
		return nil, err
	}
	result := make(map[string]trackactivities.TrackActivity, 0)
	for _, model := range models {
		result[model.SpotifyId] = model
	}
	return result, nil
}
