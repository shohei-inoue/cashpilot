package controller

import (
	"errors"
	"net/http"

	"backend/internal/apperrors"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
)

// UserController は認証済みユーザー情報用コントローラー
type UserController struct {
	userUsecase usecase.UserUsecase
}

// NewUserController は UserController を生成する
func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

// Get は GET /api/user。認証必須。
func (u *UserController) Get(c *gin.Context) {
	userID := middleware.GetUserIDFromContext(c)
	if userID == 0 {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	user, err := u.userUsecase.GetUser(c.Request.Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrUserNotFound):
			response.Error(c, http.StatusNotFound, response.CodeNotFound, response.MsgUserNotFound)
		default:
			response.Error(c, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.Success(c, user.ToResponse())
}
