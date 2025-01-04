package trackactivities

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivities"
)

func (r *repository) Get(ctx context.Context, userId uint, spotifyId string) (*trackactivities.TrackActivity, error) {
	var model trackactivities.TrackActivity
	res := r.db.Where("user_id = ?", userId).Where("spotify_id = ?", spotifyId).First(&model)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &model, nil
}
