package trackactivities

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivities"
)

func (r *repository) Create(ctx context.Context, model trackactivities.TrackActivity) error {
	return r.db.Create(&model).Error
}
