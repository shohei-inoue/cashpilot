package controller

import (
	"net/http"
	"regexp"

	"backend/internal/domain"
	"backend/internal/jwt"
	"backend/internal/repository"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost     = 10
	minPasswordLen = 8
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// AuthController は認証用コントローラー
type AuthController struct {
	jwtSecret    string
	userRepo    repository.UserRepository
}

// NewAuthController は AuthController を生成する
func NewAuthController(jwtSecret string, userRepo repository.UserRepository) *AuthController {
	return &AuthController{jwtSecret: jwtSecret, userRepo: userRepo}
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

	if !emailRegex.MatchString(req.Email) {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidEmail, response.MsgInvalidEmail)
		return
	}
	if len(req.Password) < minPasswordLen {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidPassword, response.MsgInvalidPassword)
		return
	}

	existing, _ := a.userRepo.FindByEmail(c.Request.Context(), req.Email)
	if existing != nil {
		response.Error(c, http.StatusConflict, response.CodeConflict, response.MsgEmailAlreadyExists)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, response.MsgFailedToHash)
		return
	}

	u, err := a.userRepo.Create(c.Request.Context(), req.Email, string(hash))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, response.MsgFailedToCreateUser)
		return
	}

	token, err := jwt.Issue(a.jwtSecret, u)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, response.MsgFailedToIssueToken)
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

	u, hash, err := a.userRepo.FindByEmailWithPassword(c.Request.Context(), req.Email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		return
	}
	if u == nil {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgInvalidCredentials)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)); err != nil {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgInvalidCredentials)
		return
	}

	token, err := jwt.Issue(a.jwtSecret, u)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, response.MsgFailedToIssueToken)
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
