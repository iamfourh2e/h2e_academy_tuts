package middleware

import (
	"go_tuts/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		} else {
			res, err := models.ClaimToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
			// if res.Role != "admin" {
			// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			// 	c.Abort()
			// 	return
			// }
			//context = operational, carry information
			c.Set("user", res)
		}
	}
}
