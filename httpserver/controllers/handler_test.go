package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthz(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		path       string
		want       string
		statusCode int
	}{
		{
			name:       "healthz_test01",
			method:     http.MethodGet,
			path:       "/healthz",
			want:       "ok",
			statusCode: http.StatusOK,
		},
		{
			name:       "healthz_test02",
			method:     http.MethodPost,
			path:       "/healthz",
			want:       "Method not allowed",
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, tc.path, nil)
			responseRecorder := httptest.NewRecorder()
			Healthz(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
