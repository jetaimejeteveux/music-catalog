package memberships

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Email     string `gorm:"type:varchar(100);unique;not null"`
		Username  string `gorm:"type:varchar(100);unique;not null"`
		Password  string `gorm:"type:varchar(100);not null"`
		CreatedBy string `gorm:"type:varchar(100);not null"`
		UpdatedBy string `gorm:"type:varchar(100);not null"`
	}
)

type (
	SignupRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
