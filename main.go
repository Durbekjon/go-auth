package main

import (
	"apps/go-auth/src/config"
	"apps/go-auth/src/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.DBConnect()
}

func main() {
	r := gin.Default()
	prefix := "/api/v1"
	routes.AuthRoutes(r, prefix)

	r.Run()
}
