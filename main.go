package locationapi

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// User logger Middleware
	router.Use(LoggerMiddleware())
	router.GET("/", handleGetRoot)
	router.Run(":8080")
}

/* Private*/
func handleGetRoot(c *gin.Context) {
	c.String(http.StatusOK, "Hey, from Location API")
}

/* Public */
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}
