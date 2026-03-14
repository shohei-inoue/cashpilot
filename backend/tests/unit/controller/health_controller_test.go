package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func TestHealthController_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)

	hc := controller.NewHealthController()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest(http.MethodGet, "/api/health", nil)

	hc.Get(c)

	if c.Writer.Status() != http.StatusOK {
		t.Errorf("status = %d, want %d", c.Writer.Status(), http.StatusOK)
	}
}
