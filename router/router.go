package router

import (
	"Gin_Chat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/getuserlist", service.GetUserList)
	return r
}
