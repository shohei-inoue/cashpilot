package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/router"

	"github.com/gin-gonic/gin"
)

func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.CORS())

	api := r.Group("/api")
	hc := controller.NewHealthController()
	router.Setup(api, hc)

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}
