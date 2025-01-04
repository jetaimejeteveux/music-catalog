package trackactivities

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_repository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	now := time.Now()
	isLiked := true
	type args struct {
		model trackactivities.TrackActivity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				model: trackactivities.TrackActivity{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserId:    1,
					SpotifyId: "spotifyId",
					IsLiked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "track_activities" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserId,
						args.model.SpotifyId,
						args.model.IsLiked,
						args.model.CreatedBy,
						args.model.UpdatedBy,
					).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "Error: database error",
			args: args{
				model: trackactivities.TrackActivity{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserId:    1,
					SpotifyId: "spotifyId",
					IsLiked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "track_activities" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserId,
						args.model.SpotifyId,
						args.model.IsLiked,
						args.model.CreatedBy,
						args.model.UpdatedBy,
					).WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			if err := r.Create(context.Background(), tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
