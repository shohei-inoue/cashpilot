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

// GoalController は目標用コントローラー
type GoalController struct {
	goalUsecase usecase.GoalUsecase
}

// NewGoalController は GoalController を生成する
func NewGoalController(goalUsecase usecase.GoalUsecase) *GoalController {
	return &GoalController{goalUsecase: goalUsecase}
}

// List は GET /api/goals。認証必須。
func (c *GoalController) List(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	goals, err := c.goalUsecase.List(ctx.Request.Context(), userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		return
	}

	res := make([]*domain.GoalResponse, len(goals))
	for i, g := range goals {
		res[i] = g.ToResponse()
	}
	response.Success(ctx, gin.H{"goals": res})
}

// Get は GET /api/goals/:id。認証必須。
func (c *GoalController) Get(ctx *gin.Context) {
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

	goal, err := c.goalUsecase.Get(ctx.Request.Context(), id, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrGoalNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgGoalNotFound)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.Success(ctx, goal.ToResponse())
}

// Create は POST /api/goals。認証必須。
func (c *GoalController) Create(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	var req domain.GoalCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	goal, err := c.goalUsecase.Create(ctx.Request.Context(), userID, req.Name, req.TargetAmount, req.Deadline)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidGoalName):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidGoalName)
		case errors.Is(err, apperrors.ErrInvalidTargetAmount):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidTargetAmount)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeInternalError, response.MsgInternalError)
		}
		return
	}

	response.Created(ctx, goal.ToResponse())
}

// Update は PUT /api/goals/:id。認証必須。
func (c *GoalController) Update(ctx *gin.Context) {
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

	var req domain.GoalUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	goal, err := c.goalUsecase.Update(ctx.Request.Context(), id, userID, req.Name, req.TargetAmount, req.Deadline)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrGoalNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgGoalNotFound)
		case errors.Is(err, apperrors.ErrInvalidGoalName):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidGoalName)
		case errors.Is(err, apperrors.ErrInvalidTargetAmount):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidTargetAmount)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeInternalError, response.MsgInternalError)
		}
		return
	}

	response.Success(ctx, goal.ToResponse())
}

// Delete は DELETE /api/goals/:id。認証必須。
func (c *GoalController) Delete(ctx *gin.Context) {
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

	err = c.goalUsecase.Delete(ctx.Request.Context(), id, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrGoalNotFound):
			response.Error(ctx, http.StatusNotFound, response.CodeNotFound, response.MsgGoalNotFound)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.NoContent(ctx)
}
