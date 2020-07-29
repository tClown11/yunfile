package main

import (
	"log"
	"user-client/router"
	"user-client/serviceclient"
)

func init() {
	serviceclient.RegisterService()
}

func main() {
	r := router.NewRouter()
	if err := r.Run("0.0.0.0:" + serviceclient.Port); err != nil {
		log.Print(err.Error())
	}
}
