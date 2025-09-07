package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/getters"
	"github.com/shurcooL/githubv4"
)

func LatestCommit(c *gin.Context) {
	client, exists := c.Get("github_client")
	if !exists {
		c.JSON(500, gin.H{"error": "GitHub client not found"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(500, gin.H{"error": "Username not found"})
		return
	}

	LatestCommit, err := getters.GetMostRecentCommit(client.(*githubv4.Client), username.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, LatestCommit)
}
