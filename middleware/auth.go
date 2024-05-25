package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func sha256Sum(s string) []byte {
	sum := sha256.Sum256([]byte(s))
	arr := make([]byte, len(sum))
	copy(arr, sum[:])

	return arr
}

func secureCompare(a, b string) int {
	aSum := sha256Sum(a)
	bSum := sha256Sum(b)

	return subtle.ConstantTimeCompare(aSum, bSum)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		whitelistedIPs := strings.Split(os.Getenv("WHITELISTED_IPS"), ",")
		requestedIP := strings.Split(c.Request.Header.Get("X-Forwarded-For"), ",")[0]

		whitelisted := false
		for _, value := range whitelistedIPs {
			fmt.Println(requestedIP, value)
			if secureCompare(requestedIP, value) == 1 {
				whitelisted = true
			}
		}

		if !whitelisted {
			c.AbortWithStatusJSON(403, gin.H{"error": "ACCESS DENIED"})
		}
		c.Next()
	}
}
