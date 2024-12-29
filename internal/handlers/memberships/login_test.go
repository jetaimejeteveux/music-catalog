package memberships

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)
	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       memberships.LoginResponse
		wantErr            bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@test.com",
					Password: "password",
				}).Return("access_token", nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedBody: memberships.LoginResponse{
				AccessToken: "access_token",
			},
			wantErr: false,
		},
		{
			name:               "ERROR: invalid JSON request",
			mockFn:             func() {},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       memberships.LoginResponse{},
			wantErr:            true,
		},
		{
			name: "ERROR: service returns error",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@test.com",
					Password: "password",
				}).Return("", assert.AnError)
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       memberships.LoginResponse{}, // Empty for error case
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				api,
				mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()
			endpoint := `/memberships/login`

			var body *bytes.Reader
			if tt.name == "ERROR: invalid JSON request" {
				// Invalid JSON for binding error test
				body = bytes.NewReader([]byte(`{invalid json}`))
			} else {
				reqBody := memberships.LoginRequest{
					Email:    "test@test.com",
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

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := memberships.LoginResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			} else {
				// Error case - check for error response
				var errorResponse map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)
				assert.Contains(t, errorResponse, "error")
				assert.NotEmpty(t, errorResponse["error"])
			}
		})
	}
}
