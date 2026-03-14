package controller

import (
	"net/http"
	"strings"

	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"
	"backend/internal/response"

	"github.com/gin-gonic/gin"
)

// AnalyticsController は集計用コントローラー
type AnalyticsController struct {
	analyticsUsecase usecase.AnalyticsUsecase
}

// NewAnalyticsController は AnalyticsController を生成する
func NewAnalyticsController(analyticsUsecase usecase.AnalyticsUsecase) *AnalyticsController {
	return &AnalyticsController{analyticsUsecase: analyticsUsecase}
}

func parseCashflowGroupBy(s string) repository.CashflowGroupBy {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "all":
		return repository.CashflowGroupByAll
	case "daily":
		return repository.CashflowGroupByDaily
	case "weekly":
		return repository.CashflowGroupByWeekly
	case "yearly":
		return repository.CashflowGroupByYearly
	default:
		return repository.CashflowGroupByMonthly
	}
}

// Summary は GET /api/transactions/summary。認証必須。
func (c *AnalyticsController) Summary(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	from := parseOptionalDate(ctx.Query("from"))
	to := parseOptionalDate(ctx.Query("to"))
	accountID := parseOptionalInt(ctx.Query("account_id"))
	categoryID := parseOptionalInt(ctx.Query("category_id"))

	s, err := c.analyticsUsecase.GetSummary(ctx.Request.Context(), userID, from, to, accountID, categoryID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		return
	}

	response.Success(ctx, gin.H{"summary": s})
}

// Cashflow は GET /api/analytics/cashflow。認証必須。クエリ group_by=all|daily|weekly|monthly|yearly（デフォルト: monthly）
func (c *AnalyticsController) Cashflow(ctx *gin.Context) {
	userID := middleware.GetUserIDFromContext(ctx)
	if userID == 0 {
		response.Error(ctx, http.StatusUnauthorized, response.CodeUnauthorized, response.MsgUnauthorized)
		return
	}

	groupBy := parseCashflowGroupBy(ctx.Query("group_by"))
	from := parseOptionalDate(ctx.Query("from"))
	to := parseOptionalDate(ctx.Query("to"))
	accountID := parseOptionalInt(ctx.Query("account_id"))
	categoryID := parseOptionalInt(ctx.Query("category_id"))

	items, err := c.analyticsUsecase.GetCashflow(ctx.Request.Context(), userID, groupBy, from, to, accountID, categoryID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.CodeDatabaseError, response.MsgDatabaseError)
		return
	}

	if items == nil {
		items = []*domain.CashflowByPeriod{}
	}
	response.Success(ctx, gin.H{"items": items, "group_by": string(groupBy)})
}
