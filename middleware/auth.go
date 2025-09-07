package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jasonlovesdoggo/katib/auth"
)

// Whitelisted users (will use system GH PAT, no need to provide their own). Must be lowercase.
var whitelist = []string{"jasonlovesdoggo", "araf821"}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "username parameter is required",
				"usage": "Add ?username=your_github_username to the URL",
			})
			c.Abort()
			return
		}

		// Check if user is whitelisted (will use system GH PAT)
		if slices.Contains(whitelist, strings.ToLower(username)) {
			// Use default client for whitelisted users
			c.Set("github_client", auth.Client)
			c.Set("username", username)
			c.Next()
			return
		}

		// Non-whitelisted users must provide GH_PAT via header only (security)
		ghPat := c.GetHeader("Authorization")
		if ghPat == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":         "GitHub Personal Access Token required for non-whitelisted users",
				"usage":         "Add header: Authorization: Bearer YOUR_GITHUB_PAT",
				"how_to_create": "GitHub Settings → Developer settings → Personal access tokens → Generate new token (no special permissions needed, just for rate limiting)",
			})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(ghPat) > 7 && ghPat[:7] == "Bearer " {
			ghPat = ghPat[7:]
		}

		// Create GitHub client with provided PAT
		client, err := auth.CreateClientFromPAT(ghPat)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid GitHub Personal Access Token",
			})
			c.Abort()
			return
		}

		c.Set("github_client", client)
		c.Set("username", username)
		c.Next()
	}
}
