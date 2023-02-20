package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)



func Logger() gin.HandlerFunc {

    return func(ctx *gin.Context) {
		t:= time.Now()

		
		auth := ctx.Request.Header.Get("Authorization")

		if auth == "" {
			ctx.JSON(401, &gin.H{
				"code":  401,
				"error": "Invalid Authorization header",
			})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(auth, " ", 2)

		if parts[0] != "Bearer" {
			ctx.JSON(401, &gin.H{
				"code":  401,
				"error": "Invalid Authorization header",
			})
			ctx.Abort()
			return
		}
        
        
		ctx.Set("exemple", parts[1])
		// before request
         
		ctx.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := ctx.Writer.Status()
		log.Println(status)
	}

}