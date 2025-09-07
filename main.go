package main

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/routes"
)

var cacheTime = time.Second * 60 // Cache time in seconds (this leaves me with using around 1/82nd of my rate limit)
func main() {
	store := persistence.NewInMemoryStore(time.Second)

	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(routes.NotFoundHandler)
	r.GET("/healthcheck", routes.HealthCheck)
	r.GET("/commits/latest", cache.CachePage(store, cacheTime, routes.LatestCommit))
	r.GET("/v2/commits/latest", cache.CachePage(store, cacheTime, routes.LatestCommitsV2))
	r.GET("/streak", cache.CachePage(store, cacheTime, routes.StreakInfo))
	r.Run()
}
