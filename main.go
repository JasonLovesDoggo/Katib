package main

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/middleware"
	"github.com/jasonlovesdoggo/katib/routes"
)

var cacheTime = time.Second * 120 // Cache time in seconds (this leaves me with using around 1/164 of my rate limit)
func main() {
	store := persistence.NewInMemoryStore(time.Second)

	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(routes.NotFoundHandler)

	// Public routes (no auth required)
	r.GET("/", routes.DocsHandler(DocsHTML))
	r.GET("/healthcheck", routes.HealthCheck)

	// Apply auth middleware to API routes
	api := r.Group("/")
	api.Use(middleware.AuthMiddleware())
	api.GET("/commits/latest", cache.CachePage(store, cacheTime, routes.LatestCommit))
	api.GET("/v2/commits/latest", cache.CachePage(store, cacheTime, routes.LatestCommitsV2))
	api.GET("/streak", cache.CachePage(store, cacheTime, routes.StreakInfo))

	r.Run()
}
