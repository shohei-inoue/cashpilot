package controller

import (
	"errors"
	"net/http"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
)

// SimulationController はシミュレーション用コントローラー
type SimulationController struct {
	simulationUsecase usecase.SimulationUsecase
}

// NewSimulationController は SimulationController を生成する
func NewSimulationController(simulationUsecase usecase.SimulationUsecase) *SimulationController {
	return &SimulationController{simulationUsecase: simulationUsecase}
}

// Run は POST /api/simulation/run。認証必須。
func (c *SimulationController) Run(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	var req domain.SimulationRunRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidRequest)
		return
	}

	res, err := c.simulationUsecase.Run(ctx.Request.Context(), userID, &req)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidSimulationPeriod):
			response.Error(ctx, http.StatusBadRequest, response.CodeInvalidRequest, response.MsgInvalidSimulationPeriod)
		default:
			response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		}
		return
	}

	response.Success(ctx, res)
}
