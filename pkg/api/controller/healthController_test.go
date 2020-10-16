package controller

import (
	"net/http"
	"net/http/httptest"
	"projeto-ra-api-go/pkg/api/controller/mocks"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestHealthHandler_Healthcheck(t *testing.T) {
	type fields struct {
		service *mocks.HealthChecker
	}

	tests := []struct {
		name               string
		fields             fields
		wantHttpStatusCode int
		mock               func(fs *mocks.HealthChecker)
	}{
		{name: "sucesss",
			fields: fields{
				service: new(mocks.HealthChecker),
			},

			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.HealthChecker) {
				fs.On("Check", mock.Anything, mock.Anything).Return(nil)

			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &HealthHandler{
				hc: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodGet, "/health", nil)
			recorder := httptest.NewRecorder()

			p.Healthcheck(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}
