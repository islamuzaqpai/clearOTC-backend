package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "successful health check",
			method:         http.MethodGet,
			path:           "/health",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "method not allowed",
			method:         http.MethodPost,
			path:           "/health",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "route not found",
			method:         http.MethodGet,
			path:           "/unknown",
			expectedStatus: http.StatusNotFound,
		},
	}

	router := NewRouter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}
		})
	}
}
