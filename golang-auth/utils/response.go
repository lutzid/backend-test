package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Values  interface{} `json:"values,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func WriteSuccess(ctx gin.Context, message string, data interface{}) {
	body := Message{
		Status:  "success",
		Message: message,
		Values:  data,
	}

	ctx.JSON(http.StatusOK, body)
}

func WriteError(ctx gin.Context, httpStatusCode int, message string) {
	body := Message{
		Status:  "error",
		Message: message,
	}

	ctx.JSON(httpStatusCode, body)
}
