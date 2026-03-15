package api

import (
	"backend/db"
	"backend/internal/config"
	"backend/internal/middleware"
	logicRouter "backend/internal/logic/router"

	"github.com/gin-gonic/gin"
)

const testJWTSecret = "test-jwt-secret-at-least-32-characters"

// SetupRouterWithDB は DB 接続済みのルーターを返す。接続失敗時は nil, error
func SetupRouterWithDB() (*gin.Engine, error) {
	cfg := config.Load()
	jwtSecret := cfg.JWTSecret
	if jwtSecret == "" {
		jwtSecret = testJWTSecret
	}

	gormDB, err := db.NewGormDB(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.CORS())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	logicRouter.SetupHealth(api, gormDB)
	logicRouter.SetupAuth(api, gormDB, jwtSecret)
	logicRouter.SetupUser(api, gormDB, jwtSecret)
	logicRouter.SetupAccount(api, gormDB, jwtSecret)
	logicRouter.SetupCategory(api, gormDB, jwtSecret)
	logicRouter.SetupAnalytics(api, gormDB, jwtSecret)
	logicRouter.SetupTransaction(api, gormDB, jwtSecret)
	logicRouter.SetupGoal(api, gormDB, jwtSecret)
	logicRouter.SetupSimulation(api, gormDB, jwtSecret)
	return r, nil
}
