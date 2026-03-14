package controller

import (
	"errors"
	"net/http"
	"strconv"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
)

// AccountController は口座用コントローラー
type AccountController struct {
	accountUsecase usecase.AccountUsecase
}

// NewAccountController は AccountController を生成する
func NewAccountController(accountUsecase usecase.AccountUsecase) *AccountController {
	return &AccountController{accountUsecase: accountUsecase}
}

// List は GET /api/accounts。認証必須。
func (a *AccountController) List(c *gin.Context) {
	userID := middleware.GetUserIDFromContext(c)
	if userID == 0 {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	accounts, err := a.accountUsecase.List(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		return
	}

	res := make([]*domain.AccountResponse, len(accounts))
	for i, acc := range accounts {
		res[i] = acc.ToResponse()
	}
	response.Success(c, gin.H{"accounts": res})
}

// Get は GET /api/accounts/:id。認証必須。
func (a *AccountController) Get(c *gin.Context) {
	userID := middleware.GetUserIDFromContext(c)
	if userID == 0 {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	account, err := a.accountUsecase.Get(c.Request.Context(), id, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrAccountNotFound):
			response.Error(c, http.StatusNotFound, response.CodeNotFound, response.MsgAccountNotFound)
		default:
			response.Error(c, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.Success(c, account.ToResponse())
}

// Create は POST /api/accounts。認証必須。
func (a *AccountController) Create(c *gin.Context) {
	userID := middleware.GetUserIDFromContext(c)
	if userID == 0 {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	var req domain.AccountCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	account, err := a.accountUsecase.Create(c.Request.Context(), userID, req.Type, req.Name)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidAccountType):
			response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidAccountType)
		case errors.Is(err, apperrors.ErrInvalidAccountName):
			response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidAccountName)
		default:
			response.Error(c, http.StatusInternalServerError, response.CodeInternalError, response.MsgInternalError)
		}
		return
	}

	response.Created(c, account.ToResponse())
}

// Update は PUT /api/accounts/:id。認証必須。
func (a *AccountController) Update(c *gin.Context) {
	userID := middleware.GetUserIDFromContext(c)
	if userID == 0 {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	var req domain.AccountUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	account, err := a.accountUsecase.Update(c.Request.Context(), id, userID, req.Type, req.Name)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrAccountNotFound):
			response.Error(c, http.StatusNotFound, response.CodeNotFound, response.MsgAccountNotFound)
		case errors.Is(err, apperrors.ErrInvalidAccountType):
			response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidAccountType)
		case errors.Is(err, apperrors.ErrInvalidAccountName):
			response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidAccountName)
		default:
			response.Error(c, http.StatusInternalServerError, response.CodeInternalError, response.MsgInternalError)
		}
		return
	}

	response.Success(c, account.ToResponse())
}

// Delete は DELETE /api/accounts/:id。認証必須。
func (a *AccountController) Delete(c *gin.Context) {
	userID := middleware.GetUserIDFromContext(c)
	if userID == 0 {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	err = a.accountUsecase.Delete(c.Request.Context(), id, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrAccountNotFound):
			response.Error(c, http.StatusNotFound, response.CodeNotFound, response.MsgAccountNotFound)
		default:
			response.Error(c, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.NoContent(c)
}
