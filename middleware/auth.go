package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		whitelistedIPs := strings.Split(os.Getenv("WHITELISTED_IPS"), ",")
		whitelisted := false
		for _, v := range whitelistedIPs {
			fmt.Printf("Client IP: %s, item:  %s\n", c.ClientIP(), v)
			if v == c.ClientIP() {
				whitelisted = true
			}
		}
		if !whitelisted {
			c.AbortWithStatusJSON(403, gin.H{"error": "ACCESS DENIED"})
		}
		c.Next()
	}
}
