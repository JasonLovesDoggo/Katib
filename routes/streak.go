package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/auth"
	"github.com/jasonlovesdoggo/katib/getters"
)

func StreakInfo(c *gin.Context) {
	streakInfo, err := getters.GetStreakInfo(auth.Client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, streakInfo)
}
