package tracks

import (
	"context"
	"fmt"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivites"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *service) UpsertTrackActivities(ctx context.Context, userId uint, request trackactivites.TrackActivityReqest) error {
	activity, err := s.trackActivitiesRepo.Get(ctx, userId, request.SpotifyId)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error getting track activities")
		return err
	}
	if err == gorm.ErrRecordNotFound || activity == nil {
		err = s.trackActivitiesRepo.Create(ctx, trackactivites.TrackActivity{
			UserId:    userId,
			SpotifyId: request.SpotifyId,
			IsLiked:   request.IsLiked,
			CreatedBy: fmt.Sprintf("%d", userId),
			UpdatedBy: fmt.Sprintf("%d", userId),
		})
		if err != nil {
			log.Error().Err(err).Msg("error creating track activities")
			return err
		}
		return nil
	}

	activity.IsLiked = request.IsLiked
	err = s.trackActivitiesRepo.Update(ctx, *activity)
	if err != nil {
		log.Error().Err(err).Msg("error updating track activities")
		return err
	}
	return nil
}
