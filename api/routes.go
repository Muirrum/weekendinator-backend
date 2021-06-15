package api

import (
	"github.com/Muirrum/weekendinator-backend/api/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) error {
	auth_rg := r.Group("/auth")
	auth.SetupRoutes(auth_rg)

	return nil
}
