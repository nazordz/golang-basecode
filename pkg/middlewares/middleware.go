package middlewares

import (
	"net/http"

	"github.com/devgoorita/golang-basecode/internal/utils"
	"github.com/devgoorita/golang-basecode/pkg"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var Gzipping = gzip.Gzip(gzip.BestCompression)

var CorsConfig = cors.New(cors.Config{
	AllowAllOrigins: true,
	AllowMethods:    []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
	AllowHeaders:    []string{"Content-Type", "Authorization", "Accept-Encoding"},
})

func GinMode() string {
	var mode string
	if pkg.GodotEnv("GO_ENV") != "development" {
		mode = gin.ReleaseMode
	} else {
		mode = gin.DebugMode
	}
	return mode
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
