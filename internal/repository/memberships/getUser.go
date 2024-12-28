package memberships

import "github.com/jetaimejeteveux/music-catalog/internal/models/memberships"

func (r *repository) GetUser(email, username string, id uint) (*memberships.User, error) {
	user := memberships.User{}
	tx := r.db.Where("email = ?", email).Or("username =? ", username).Or("id =? ", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
