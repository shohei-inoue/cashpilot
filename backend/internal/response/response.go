package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorBody はエラーレスポンスのボディ
type ErrorBody struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Error はエラーレスポンスを返す
func Error(c *gin.Context, status int, code, message string) {
	c.JSON(status, ErrorBody{
		Error: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    code,
			Message: message,
		},
	})
}

// Success は 200 OK でデータを返す
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Created は 201 Created でデータを返す
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// NoContent は 204 No Content を返す
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
