package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimit(limiter *ratelimit.Bucket) gin.HandlerFunc {
	fmt.Println("haha")
	return func(c *gin.Context) {
		fmt.Println("haha2")
		if limiter.TakeAvailable(1) == 0 {
			c.String(http.StatusTooManyRequests, fmt.Sprintf("Rate limit exceeded. Try again in %v seconds.", limiter.Available()))
			c.Abort()
			return
		}
		c.Next()
	}
}
