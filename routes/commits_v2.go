package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/getters"
	"github.com/shurcooL/githubv4"
)

func LatestCommitsV2(c *gin.Context) {
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

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	response, err := getters.GetCommitsList(client.(*githubv4.Client), username.(string), limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}
