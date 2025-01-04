package trackactivities

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivites"
)

func (r *repository) Get(ctx context.Context, userId uint, spotifyId string) (*trackactivites.TrackActivity, error) {
	var model trackactivites.TrackActivity
	res := r.db.Where("user_id = ?", userId).Where("spotify_id = ?", spotifyId).First(&model)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &model, nil
}
