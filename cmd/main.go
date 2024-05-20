package main

import (
	cicdRouter "github.com/JustinHung0407/cicdemo/pkg/router"
	"log"
)

func main() {
	router := cicdRouter.SetupRouter()
	err := router.Run("0.0.0.0:8085")
	if err != nil {
		log.Fatalf("Error while starting server: " + err.Error())
	}
}
