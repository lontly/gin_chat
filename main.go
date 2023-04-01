package main

import (
	"Gin_Chat/router"
	"Gin_Chat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := router.Router()
	r.Run()
}
