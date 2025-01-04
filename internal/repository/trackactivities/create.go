package trackactivities

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivites"
)

func (r *repository) Create(ctx context.Context, model trackactivites.TrackActivity) error {
	return r.db.Create(&model).Error
}
