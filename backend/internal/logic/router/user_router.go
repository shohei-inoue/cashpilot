package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupUser はユーザールートを登録する
func SetupUser(r *gin.RouterGroup, db *gorm.DB, jwtSecret string) {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	uc := controller.NewUserController(userUsecase)

	r.GET("/user", middleware.Auth(jwtSecret), uc.Get)
}
