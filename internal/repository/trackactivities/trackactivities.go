package trackactivities

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivites"
)

func (r *repository) Create(ctx context.Context, model trackactivites.TrackActivity) error {
	return r.db.Create(&model).Error
}

func (r *repository) Update(ctx context.Context, model trackactivites.TrackActivity) error {
	return r.db.Save(&model).Error
}

func (r *repository) Get(ctx context.Context, userId uint, spotifyId string) (*trackactivites.TrackActivity, error) {
	var model trackactivites.TrackActivity
	res := r.db.Where("user_id = ?", userId).Where("spotify_id = ?", spotifyId).First(&model)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *repository) GetBulkSpotifyIDs(ctx context.Context, userId uint, spotifyIds []string) (map[string]trackactivites.TrackActivity, error) {
	var models []trackactivites.TrackActivity
	res := r.db.Where("user_id = ?", userId).Where("spotify_id IN (?)", spotifyIds).Find(&models)
	if err := res.Error; err != nil {
		return nil, err
	}
	result := make(map[string]trackactivites.TrackActivity, 0)
	for _, model := range models {
		result[model.SpotifyId] = model
	}
	return result, nil
}
