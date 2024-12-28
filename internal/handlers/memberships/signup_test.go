package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Signup(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)
	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().Signup(memberships.SignupRequest{
					Email:    "test@test.com",
					Username: "testuser",
					Password: "password",
				}).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "ERROR: username or email exists",
			mockFn: func() {
				mockSvc.EXPECT().Signup(memberships.SignupRequest{
					Email:    "test@test.com",
					Username: "testuser",
					Password: "password",
				}).Return(errors.New("username or email exists"))
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "ERROR: invalid JSON request",
			mockFn: func() {

			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()
			endpoint := `/memberships/signup`

			var body *bytes.Reader
			if tt.name == "ERROR: invalid JSON request" {
				// Invalid JSON for binding error test
				body = bytes.NewReader([]byte(`{invalid json}`))
			} else {
				reqBody := memberships.SignupRequest{
					Email:    "test@test.com",
					Username: "testuser",
					Password: "password",
				}
				val, err := json.Marshal(reqBody)
				assert.NoError(t, err)
				body = bytes.NewReader(val)
			}

			request, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			h.ServeHTTP(w, request)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
