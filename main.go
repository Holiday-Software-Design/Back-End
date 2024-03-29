package main

import (
	"hr/configs/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Init(r)
	err := r.Run(":3000")
	if err != nil {
		log.Fatal("Server start error", err)
	}

}
