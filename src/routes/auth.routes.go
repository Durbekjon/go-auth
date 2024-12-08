package routes

import (
	"apps/go-auth/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, prefix string) {
	r.POST(prefix+"/auth/register", controllers.Register)
	r.POST(prefix+"/auth/login", controllers.Login)
	r.POST(prefix+"/auth/refresh", controllers.RefreshTokens)
}
