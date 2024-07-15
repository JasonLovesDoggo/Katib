package routes

import (
	"github.com/gin-gonic/gin"
	"katib/auth"
	"katib/getters"
)

func LatestCommit(c *gin.Context) {

	LatestCommit, err := getters.GetMostRecentCommit(auth.Client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, LatestCommit)
}
