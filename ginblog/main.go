package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	model.InitMysql()
	routes.InitRouter()
}
