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

// CategoryController はカテゴリ用コントローラー
type CategoryController struct {
	categoryUsecase usecase.CategoryUsecase
}

// NewCategoryController は CategoryController を生成する
func NewCategoryController(categoryUsecase usecase.CategoryUsecase) *CategoryController {
	return &CategoryController{categoryUsecase: categoryUsecase}
}

// List は GET /api/categories。認証必須。
func (c *CategoryController) List(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	categories, err := c.categoryUsecase.List(ctx.Request.Context(), userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		return
	}

	res := make([]*domain.CategoryResponse, len(categories))
	for i, cat := range categories {
		res[i] = cat.ToResponse()
	}
	response.Success(ctx, gin.H{"categories": res})
}

// Get は GET /api/categories/:id。認証必須。
func (c *CategoryController) Get(ctx *gin.Context) {
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

	category, err := c.categoryUsecase.Get(ctx.Request.Context(), id, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrCategoryNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgCategoryNotFound)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.Success(ctx, category.ToResponse())
}

// Create は POST /api/categories。認証必須。
func (c *CategoryController) Create(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	var req domain.CategoryCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	category, err := c.categoryUsecase.Create(ctx.Request.Context(), userID, req.Type, req.Name)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidCategoryType):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidCategoryType)
		case errors.Is(err, apperrors.ErrInvalidCategoryName):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidCategoryName)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeInternalError, response.MsgInternalError)
		}
		return
	}

	response.Created(ctx, category.ToResponse())
}

// Update は PUT /api/categories/:id。認証必須。
func (c *CategoryController) Update(ctx *gin.Context) {
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

	var req domain.CategoryUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	category, err := c.categoryUsecase.Update(ctx.Request.Context(), id, userID, req.Type, req.Name)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrCategoryNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgCategoryNotFound)
		case errors.Is(err, apperrors.ErrInvalidCategoryType):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidCategoryType)
		case errors.Is(err, apperrors.ErrInvalidCategoryName):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidCategoryName)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeInternalError, response.MsgInternalError)
		}
		return
	}

	response.Success(ctx, category.ToResponse())
}

// Delete は DELETE /api/categories/:id。認証必須。
func (c *CategoryController) Delete(ctx *gin.Context) {
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

	err = c.categoryUsecase.Delete(ctx.Request.Context(), id, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrCategoryNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgCategoryNotFound)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.NoContent(ctx)
}
