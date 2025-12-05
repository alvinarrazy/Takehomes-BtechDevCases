package util

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Message string `json:"message,omitempty"`
	Payload any    `json:"data,omitempty"`
}

func returnJSONResponse(w gin.ResponseWriter, message string, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(JSONResponse{
		Message: message,
		Payload: data,
	})

}

func SuccessResponse(w gin.ResponseWriter, data any, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	returnJSONResponse(w, "success", data, code)
}

func ErrorResponse(w gin.ResponseWriter, message string, statusCode ...int) {
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	returnJSONResponse(w, message, nil, code)
}
