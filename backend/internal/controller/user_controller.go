package controller

import (
	"net/http"

	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
)

// UserController は認証済みユーザー情報用コントローラー
type UserController struct {
	userRepo repository.UserRepository
}

// NewUserController は UserController を生成する
func NewUserController(userRepo repository.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

// Get は GET /api/user。認証必須。
func (u *UserController) Get(c *gin.Context) {
	userID := middleware.GetUserIDFromContext(c)
	if userID == 0 {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	user, err := u.userRepo.FindByID(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		return
	}
	if user == nil {
		response.Error(c, http.StatusNotFound, response.CodeNotFound, response.MsgUserNotFound)
		return
	}

	response.Success(c, user.ToResponse())
}
