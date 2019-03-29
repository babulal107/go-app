package main

import (
	"bitbucket.org/babulal107/go-app/app"
	"bitbucket.org/babulal107/go-app/config"
	"bitbucket.org/babulal107/go-app/helper"
	"log"
)
 
func main() {
	log.Println("Load configuration")
	configObj := config.GetConfig()
	log.Println("Start Application")
	appObj := new(app.App)
	appObj.Initialize(configObj)
	defer helper.Close(appObj)
	appObj.Run(":5000")
}
