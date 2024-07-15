package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"katib/routes"
	"time"
)

var cacheTime = time.Minute * 3 // Cache time in minutes todo: calculate the exact amount needed based off of my quota
func main() {
	store := persistence.NewInMemoryStore(time.Second)

	// Use client...

	r := gin.Default()
	r.GET("/commits/latest", cache.CachePage(store, cacheTime, routes.LatestCommit))
	r.Run()
}
