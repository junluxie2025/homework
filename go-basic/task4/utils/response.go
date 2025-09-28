package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// Unauthorized 401错误
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

// BadRequest 400错误
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

// StatusInternalServerError 500错误
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}

// NotFound 404错误
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// Forbidden 403错误
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message)
}
