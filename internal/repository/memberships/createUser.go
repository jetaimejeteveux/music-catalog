package memberships

import (
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
)

func (r *repository) CreateUser(model memberships.User) error {
	return r.db.Create(&model).Error
}
