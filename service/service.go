package service

import (
	"Gin_Chat/models"

	"github.com/gin-gonic/gin"
)

func  GetIndex(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"welcome",
	})
}

func GetUserList(c *gin.Context){
	userlist := models.GetUserList()
	c.JSON(200,gin.H{
		"message":userlist,
	})
}
