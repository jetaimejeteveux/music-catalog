package memberships

import (
	"github.com/jetaimejeteveux/music-catalog/internal/configs"
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"testing"
)

const TestPassword = "test_password_123"

func generateTestHash(t *testing.T) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(TestPassword), bcrypt.MinCost)
	if err != nil {
		t.Fatal("Failed to hash password for test:", err)
	}
	return string(hashedPassword)
}

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	type args struct {
		req memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		config  *configs.Config
		mockFn  func(args args)
	}{
		{
			name: "SUCCESS",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@test.com",
					Password: TestPassword,
				},
			},
			config: &configs.Config{
				Service: configs.Service{
					SecretJWT: "secret",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Email:    "test@test.com",
					Password: generateTestHash(t),
					Username: "testusername",
				}, nil)

			},
		},
		{
			name: "ERROR: user not found",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@test.com",
					Password: TestPassword,
				},
			},
			config:  &configs.Config{},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(nil, assert.AnError)

			},
		},
		{
			name: "ERROR: invalid password",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@test.com",
					Password: strings.Repeat("a", 100),
				},
			},
			wantErr: true,
			config:  &configs.Config{},
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Email:    "test@test.com",
					Password: generateTestHash(t),
					Username: "testusername",
				}, nil)

			},
		},
		{
			name: "ERROR: jwt creation fail",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@test.com",
					Password: strings.Repeat("a", 100),
				},
			},
			wantErr: true,
			config:  &configs.Config{},
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Email:    "test@test.com",
					Password: generateTestHash(t),
					Username: "testusername",
				}, nil)

			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cfg:        tt.config,
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
