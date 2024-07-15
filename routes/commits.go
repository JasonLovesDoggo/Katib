package routes

import (
	"github.com/gin-gonic/gin"
)

func LatestCommit(c *gin.Context) {

	LatestCommit := GetMostRecentCommit(Client)
	c.JSON(200, LatestCommit)
}
