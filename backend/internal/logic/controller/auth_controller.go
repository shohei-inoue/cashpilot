package controller

import (
	"errors"
	"net/http"

	"backend/internal/apperrors"
	"backend/internal/jwt"
	"backend/internal/logic/domain"
	"backend/internal/logic/usecase"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
)

// AuthController は認証用コントローラー
type AuthController struct {
	authUsecase usecase.AuthUsecase
}

// NewAuthController は AuthController を生成する
func NewAuthController(authUsecase usecase.AuthUsecase) *AuthController {
	return &AuthController{authUsecase: authUsecase}
}

// setTokenCookie は JWT を Cookie にセットする
func setTokenCookie(c *gin.Context, token string) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		jwt.CookieName,
		token,
		86400, // 24h
		"/",
		"",
		false, // Secure: 本番では true
		true,  // HttpOnly
	)
}

// clearTokenCookie は Cookie を削除する
func clearTokenCookie(c *gin.Context) {
	c.SetCookie(jwt.CookieName, "", -1, "/", "", false, true)
}

// Signup は POST /api/auth/signup
func (a *AuthController) Signup(c *gin.Context) {
	var req domain.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	u, token, err := a.authUsecase.Signup(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidEmail):
			response.Error(c, http.StatusBadRequest, response.CodeInvalidEmail, response.MsgInvalidEmail)
		case errors.Is(err, apperrors.ErrInvalidPassword):
			response.Error(c, http.StatusBadRequest, response.CodeInvalidPassword, response.MsgInvalidPassword)
		case errors.Is(err, apperrors.ErrEmailExists):
			response.Error(c, http.StatusConflict, response.CodeConflict, response.MsgEmailAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, response.CodeInternalError, response.MsgFailedToCreateUser)
		}
		return
	}

	setTokenCookie(c, token)
	response.Created(c, u.ToResponse())
}

// Login は POST /api/auth/login
func (a *AuthController) Login(c *gin.Context) {
	var req domain.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	u, token, err := a.authUsecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidCredentials):
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgInvalidCredentials)
		default:
			response.Error(c, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	setTokenCookie(c, token)
	response.Success(c, u.ToResponse())
}

// Logout は POST /api/auth/logout
func (a *AuthController) Logout(c *gin.Context) {
	clearTokenCookie(c)
	response.Success(c, gin.H{response.MsgStatus: response.MsgOK})
}
