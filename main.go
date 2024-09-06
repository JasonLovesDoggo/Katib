package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/routes"
	"time"
)

var cacheTime = time.Second * 30 // Cache time in seconds (this leaves me with using around 1/41st of my rate limit, assuming 1 credit per page)
func main() {
	store := persistence.NewInMemoryStore(time.Second)

	// Use client...

	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(routes.NotFoundHandler)
	r.GET("/healthcheck", routes.HealthCheck)
	r.GET("/commits/latest", cache.CachePage(store, cacheTime, routes.LatestCommit))
	r.GET("/stats/lifetime", cache.CachePage(store, time.Hour*24, routes.TotalStats)) // Update Daily
	r.Run()
}
