package middleware

import (
	"fmt"
	"os"
	"strings"
	"tech-buzzword-service/util"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		whitelistedIPs := strings.Split(os.Getenv("WHITELISTED_IPS"), ",")
		requestedIP := strings.Split(c.Request.Header.Get("X-Forwarded-For"), ",")[0]

		whitelisted := false
		for _, value := range whitelistedIPs {
			fmt.Println(requestedIP, value)
			if util.SecureCompare(requestedIP, value) == 1 {
				whitelisted = true
			}
		}

		if !whitelisted {
			c.AbortWithStatusJSON(403, gin.H{"error": "ACCESS DENIED"})
		}
		c.Next()
	}
}
