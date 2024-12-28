package memberships

import (
	"database/sql"
	"github.com/jetaimejeteveux/music-catalog/internal/configs"
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
)

func Test_service_Signup(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	type args struct {
		request memberships.SignupRequest
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
				memberships.SignupRequest{
					Email:    "test@test.com",
					Username: "testuser",
					Password: "test123",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, sql.ErrNoRows)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "ERROR: get user throws error",
			args: args{
				memberships.SignupRequest{
					Email:    "test@test.com",
					Username: "testuser",
					Password: "test123",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "ERROR: user already exists",
			args: args{
				memberships.SignupRequest{
					Email:    "test@test.com",
					Username: "testuser",
					Password: "test123",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				existingUser := memberships.User{
					Email:    "test@test.com",
					Username: "testuser",
					Password: "test123",
				}
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(&existingUser, nil)
			},
		},
		{
			name: "ERROR: bcrypt hash generation fails",
			args: args{
				memberships.SignupRequest{
					Email:    "test@test.com",
					Username: "testuser",
					Password: strings.Repeat("a", 73),
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, sql.ErrNoRows)
				// No need to expect CreateUser since bcrypt will fail before reaching that point
			},
		},
		{
			name: "ERROR: create user throws error",
			args: args{
				memberships.SignupRequest{
					Email:    "test@test.com",
					Username: "testuser",
					Password: "test123",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, sql.ErrNoRows)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.Signup(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
