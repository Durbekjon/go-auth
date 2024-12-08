package main

import (
	"apps/go-auth/src/config"
	"apps/go-auth/src/models"
)

func init() {
	config.LoadEnv()
	config.DBConnect()
}
func main() {
	config.DB.AutoMigrate(&models.User{})
}
