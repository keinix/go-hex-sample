package middlewear

import "github.com/gin-gonic/gin"

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			// don't overwrite current HTTP status code
			c.JSON(-1, c.Errors)
		}
	}
}
