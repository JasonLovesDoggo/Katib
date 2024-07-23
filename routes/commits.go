package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/auth"
	"github.com/jasonlovesdoggo/katib/getters"
)

func LatestCommit(c *gin.Context) {

	LatestCommit, err := getters.GetMostRecentCommit(auth.Client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, LatestCommit)
}
