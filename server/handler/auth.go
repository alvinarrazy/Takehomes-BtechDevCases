package handler

import (
	"net/http"
	"server/middleware"
	"server/service"
	"server/util"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		util.ErrorResponse(c.Writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	tokenResponse, err := handler.service.Login(request.Email, request.Password)
	if err != nil {
		util.ErrorResponse(c.Writer, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	setCookie(c, tokenResponse.Token)

	util.SuccessResponse(c.Writer, gin.H{
		"user": tokenResponse.Email,
	}, http.StatusOK)
}

func (handler *AuthHandler) Logout(c *gin.Context) {
	setCookie(c, "")

	util.SuccessResponse(c.Writer, gin.H{
		"message": "Logout success!",
	}, http.StatusOK)
}

func (handler *AuthHandler) GetUser(c *gin.Context) {
	tokenString := middleware.ExtractTokenFromCookie(c)
	claims, err := middleware.ValidateJWT(tokenString)
	if err != nil {
		util.ErrorResponse(c.Writer, "Invalid token", http.StatusUnauthorized)
		return
	}

	util.SuccessResponse(c.Writer, gin.H{
		"email": (*claims)["email"],
	}, http.StatusOK)
}

func setCookie(c *gin.Context, token string) {
	expires := time.Now().Add(15 * time.Minute)

	if token == "" {
		expires = time.Now().Add(-time.Hour) // Expire cookie for logout
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		Path:     "/",
		Expires:  expires,
	})
}
