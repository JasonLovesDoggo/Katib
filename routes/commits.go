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

func TotalStats(c *gin.Context) {
	LifetimeStats, err := getters.GetTotalStats(auth.Client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, LifetimeStats)

}

func GetStreak(c *gin.Context) {
	StartTime := c.Query("start")
	EndTime := c.Query("end")



	Streak, err := getters.GetStreak(auth.Client, StartTime, EndTime)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, Streak)
}


func GetLifetimeStreak(c *gin.Context) {
	LifetimeStreak, err := getters.Str(auth.Client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, LifetimeStreak)
}
