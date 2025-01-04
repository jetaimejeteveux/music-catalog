package trackactivities

import "gorm.io/gorm"

type TrackActivity struct {
	gorm.Model
	UserId    uint   `gorm:"not null"`
	SpotifyId string `gorm:"not null"`
	IsLiked   *bool
	CreatedBy string `gorm:"not null"`
	UpdatedBy string `gorm:"not null"`
}

type TrackActivityReqest struct {
	SpotifyId string `json:"spotifyID"`
	IsLiked   *bool  `json:"isLiked"`
}
