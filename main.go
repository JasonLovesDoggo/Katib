package main

import (
	"github.com/gin-gonic/gin"
	"katib/routes"
)

func main() {
	// Use client...
	r := gin.Default()
	r.GET("/commits/latest", routes.LatestCommit)
	r.Run()
}
