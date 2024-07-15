package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"katib/routes"
	"time"
)

var cacheTime = time.Second * 30 // Cache time in seconds (this leaves me with using around 1/41st of my rate limit)
func main() {
	store := persistence.NewInMemoryStore(time.Second)

	// Use client...

	r := gin.Default()
	r.NoRoute(routes.NotFoundHandler)
	r.GET("/healthcheck", routes.HealthCheck)
	r.GET("/commits/latest", cache.CachePage(store, cacheTime, routes.LatestCommit))
	r.Run()
}
