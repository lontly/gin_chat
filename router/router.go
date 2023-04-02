package router

import (
	"Gin_Chat/docs"
	"Gin_Chat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	r.GET("/user/getuserlist", service.GetUserList)
	r.GET("/user/register", service.CreateUser)
	r.GET("/user/deleteuser", service.DeleteUser)
	r.POST("/user/updateuser", service.UpdateUser)
	return r
}
