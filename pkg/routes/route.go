package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/chrisge4/mad-demo-http-service/pkg/api"
	"github.com/chrisge4/mad-demo-http-service/pkg/config"
)

func Routes(e *gin.Engine, sc *config.ServerConfig) {

	v1 := e.Group("/api/v1")
	v1.GET("/todo/:category", api.ListFn(sc))
	v1.POST("/todo/add", api.AddFn(sc))
}
