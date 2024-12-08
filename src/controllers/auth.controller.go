package controllers

import (
	"apps/go-auth/src/config"
	"apps/go-auth/src/models"
	"apps/go-auth/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	hashedPassword := utils.HashPassword(req.Password)
	user := models.User{
		ID:        utils.GenerateUUID(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	response := gin.H{
		"user": gin.H{
			"id":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
		},
		"tokens": GenerateTokens(user.ID),
	}

	utils.RespondWithJSON(c, http.StatusCreated, response)
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid password", nil)
		return
	}

	response := gin.H{
		"user": gin.H{
			"id":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
		},
		"tokens": GenerateTokens(user.ID),
	}

	utils.RespondWithJSON(c, http.StatusOK, response)
}

func RefreshTokens(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	if !ValidateRefreshToken(req.RefreshToken) {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid refresh token", nil)
		return
	}

	// Extract user ID from the refresh token
	claims, err := ExtractTokenClaims(req.RefreshToken)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token claims", err)
		return
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid user ID in token", nil)
		return
	}

	response := gin.H{
		"tokens": GenerateTokens(userId),
	}
	utils.RespondWithJSON(c, http.StatusOK, response)
}
