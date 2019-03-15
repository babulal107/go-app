package main

import (
	"bitbucket.org/go-app/app"
	"bitbucket.org/go-app/config"
	"log"
)
 
func main() {
	log.Println("Load configuration")
	configObj := config.GetConfig()
	log.Println("Start Application")
	appObj := &app.App{}
	appObj.Initialize(configObj)
	appObj.Run(":5000")
}
