package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/auth"
	"github.com/jasonlovesdoggo/katib/getters"
	"net/http"
	"time"
)

func LatestCommit(c *gin.Context) {

	Commit, err := getters.GetMostRecentCommit(auth.Client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, Commit)
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
	const layout string = "02-01-2006" // dd-mm-yyyy
	StartTime, err := time.Parse(layout, c.Param("start"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date, please use dd-mm-yyyy for start and end dates"})
	}

	Endtime, err := time.Parse(layout, c.Param("end"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date, please use dd-mm-yyyy for start and end dates"})
	}

	Streak, err := getters.GetStreak(auth.Client, StartTime, Endtime)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, Streak)
}

func GetLifetimeStreak(c *gin.Context) {
	LifetimeStreak, err := getters.GetLifetimeStreak(auth.Client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()}) // Internal server error todo: better handling
		return
	}
	c.JSON(200, LifetimeStreak)
}
