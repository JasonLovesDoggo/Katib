package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/auth"
	"github.com/jasonlovesdoggo/katib/getters"
	"strconv"
)

func LatestCommitsV2(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	response, err := getters.GetCommitsList(auth.Client, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}
