package main

import (
	"com.webapp/webapp/controller"
	"com.webapp/webapp/model"
)

func main() {
	model.Init()
	controller.Start()

}
