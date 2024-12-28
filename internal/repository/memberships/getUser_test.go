package memberships

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_repository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	now := time.Now()

	type args struct {
		email    string
		username string
		id       uint
	}
	tests := []struct {
		name    string
		args    args
		want    *memberships.User
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "SUCCESS",
			args: args{
				email:    "test@test.com",
				username: "test",
			},
			want: &memberships.User{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				Email:     "test@test.com",
				Username:  "testuser",
				Password:  "password",
				CreatedBy: "test@test.com",
				UpdatedBy: "test@test.com",
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" .+`).WithArgs(args.email, args.username, args.id, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "username", "password", "created_by", "updated_by"}).
						AddRow(1, now, now, "test@test.com", "testuser", "password", "test@test.com", "test@test.com"))
			},
		},
		{
			name: "ERROR",
			args: args{
				email:    "test@test.com",
				username: "test",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" .+`).WithArgs(args.email, args.username, args.id, 1).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			got, err := r.GetUser(tt.args.email, tt.args.username, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "GetUser(%v, %v, %v)", tt.args.email, tt.args.username, tt.args.id)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
