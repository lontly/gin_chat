package main

import (
	"Gin_Chat/router"
	"Gin_Chat/utils"
)


//main包启动app
func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := router.Router()
	r.Run()
}
