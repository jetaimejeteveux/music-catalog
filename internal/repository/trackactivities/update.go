package trackactivities

import (
	"context"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivites"
)

func (r *repository) Update(ctx context.Context, model trackactivites.TrackActivity) error {
	return r.db.Save(&model).Error
}
