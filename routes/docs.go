package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocsHandler(docsHTML []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", docsHTML)
	}
}
