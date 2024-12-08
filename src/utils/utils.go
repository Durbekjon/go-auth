package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var AccsessSecretKey = "secret"
var RefreshSecretKey = "secret"

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUUID() string {
	return uuid.New().String()
}

func RespondWithError(c *gin.Context, status int, message string, err error) {
	response := gin.H{"error": message}
	if err != nil {
		response["details"] = err.Error()
	}
	c.JSON(status, response)
}

func RespondWithJSON(c *gin.Context, status int, payload gin.H) {
	c.JSON(status, payload)
}
