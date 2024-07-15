package routes

import "github.com/gin-gonic/gin"

func NotFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"error": "Not Found - please visit https://github.com/JasonLovesDoggo/Katib for more information"})
}
