package controller

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
)

// TransactionController は取引用コントローラー
type TransactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

// NewTransactionController は TransactionController を生成する
func NewTransactionController(transactionUsecase usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{transactionUsecase: transactionUsecase}
}

func parseOptionalDate(s string) *time.Time {
	if s == "" {
		return nil
	}
	for _, layout := range []string{time.RFC3339, "2006-01-02T15:04:05", "2006-01-02"} {
		if t, err := time.Parse(layout, s); err == nil {
			return &t
		}
	}
	return nil
}

func parseOptionalInt(s string) *int {
	if s == "" {
		return nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &n
}

// List は GET /api/transactions。認証必須。
func (c *TransactionController) List(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	filter := repository.ListFilter{
		From:       parseOptionalDate(ctx.Query("from")),
		To:         parseOptionalDate(ctx.Query("to")),
		AccountID:  parseOptionalInt(ctx.Query("account_id")),
		CategoryID: parseOptionalInt(ctx.Query("category_id")),
		Limit:      50,
		Offset:     0,
	}
	if v := ctx.Query("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 100 {
			filter.Limit = n
		}
	}
	if v := ctx.Query("offset"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			filter.Offset = n
		}
	}

	transactions, err := c.transactionUsecase.List(ctx.Request.Context(), userID, filter)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		return
	}

	res := make([]*domain.TransactionResponse, len(transactions))
	for i, tx := range transactions {
		res[i] = tx.ToResponse()
	}
	response.Success(ctx, gin.H{"transactions": res})
}

// Get は GET /api/transactions/:id。認証必須。
func (c *TransactionController) Get(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	transaction, err := c.transactionUsecase.Get(ctx.Request.Context(), id, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrTransactionNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgTransactionNotFound)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.Success(ctx, transaction.ToResponse())
}

// Create は POST /api/transactions。認証必須。
func (c *TransactionController) Create(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	var req domain.TransactionCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	occurredAt := parseOptionalDate(req.OccurredAt)
	if occurredAt == nil {
		now := time.Now()
		occurredAt = &now
	}

	transaction, err := c.transactionUsecase.Create(ctx.Request.Context(), userID, req.AccountID, req.CategoryID, req.Amount, req.Memo, *occurredAt)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidAmount):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidAmount)
		case errors.Is(err, apperrors.ErrInvalidAccountRef):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidAccountRef)
		case errors.Is(err, apperrors.ErrInvalidCategoryRef):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidCategoryRef)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeInternalError, response.MsgInternalError)
		}
		return
	}

	response.Created(ctx, transaction.ToResponse())
}

// Update は PUT /api/transactions/:id。認証必須。
func (c *TransactionController) Update(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	var req domain.TransactionUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	occurredAt := parseOptionalDate(req.OccurredAt)
	if occurredAt == nil {
		now := time.Now()
		occurredAt = &now
	}

	transaction, err := c.transactionUsecase.Update(ctx.Request.Context(), id, userID, req.AccountID, req.CategoryID, req.Amount, req.Memo, *occurredAt)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrTransactionNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgTransactionNotFound)
		case errors.Is(err, apperrors.ErrInvalidAmount):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidAmount)
		case errors.Is(err, apperrors.ErrInvalidAccountRef):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidAccountRef)
		case errors.Is(err, apperrors.ErrInvalidCategoryRef):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidCategoryRef)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeInternalError, response.MsgInternalError)
		}
		return
	}

	response.Success(ctx, transaction.ToResponse())
}

// Delete は DELETE /api/transactions/:id。認証必須。
func (c *TransactionController) Delete(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	err = c.transactionUsecase.Delete(ctx.Request.Context(), id, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrTransactionNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgTransactionNotFound)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.NoContent(ctx)
}
