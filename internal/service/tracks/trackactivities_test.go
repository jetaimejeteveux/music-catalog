package tracks

import (
	"context"
	"fmt"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivities"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

func Test_service_UpsertTrackActivities(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTrackActivityRepo := NewMocktrackActivitiesRepository(mockCtrl)

	isLikedTrue := true
	isLikedFalse := false
	type args struct {
		userID  uint
		request trackactivities.TrackActivityRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success: create",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivityRequest{
					SpotifyId: "spotifyID",
					IsLiked:   &isLikedTrue,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyId).Return(nil, gorm.ErrRecordNotFound)

				mockTrackActivityRepo.EXPECT().Create(gomock.Any(), trackactivities.TrackActivity{
					UserId:    args.userID,
					SpotifyId: args.request.SpotifyId,
					IsLiked:   args.request.IsLiked,
					CreatedBy: fmt.Sprintf("%d", args.userID),
					UpdatedBy: fmt.Sprintf("%d", args.userID),
				}).Return(nil)
			},
		},
		{
			name: "success: update",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivityRequest{
					SpotifyId: "spotifyID",
					IsLiked:   &isLikedTrue,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyId).Return(&trackactivities.TrackActivity{
					IsLiked: &isLikedFalse,
				}, nil)

				mockTrackActivityRepo.EXPECT().Update(gomock.Any(), trackactivities.TrackActivity{
					IsLiked: args.request.IsLiked,
				}).Return(nil)
			},
		},

		{
			name: "failed",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivityRequest{
					SpotifyId: "spotifyID",
					IsLiked:   &isLikedTrue,
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyId).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				trackActivitiesRepo: mockTrackActivityRepo,
			}
			if err := s.UpsertTrackActivities(context.Background(), tt.args.userID, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.UpsertTrackActivities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
