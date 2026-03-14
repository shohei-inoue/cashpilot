package router

import (
	"backend/internal/logic/controller"
	"backend/internal/logic/repository"
	"backend/internal/logic/usecase"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupTransaction は取引ルートを登録する
func SetupTransaction(r *gin.RouterGroup, pool *pgxpool.Pool, jwtSecret string) {
	transactionRepo := repository.NewTransactionRepository(pool)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)
	tc := controller.NewTransactionController(transactionUsecase)

	transactions := r.Group("/transactions")
	transactions.Use(middleware.Auth(jwtSecret))
	{
		transactions.GET("", tc.List)
		transactions.GET("/:id", tc.Get)
		transactions.POST("", tc.Create)
		transactions.PUT("/:id", tc.Update)
		transactions.DELETE("/:id", tc.Delete)
	}
}
