package main
 
import (
	"bitbucket.org/go-api/app"
	"bitbucket.org/go-api/config"
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
