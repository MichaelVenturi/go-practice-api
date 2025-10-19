package middlewares

import (
	"net/http"

	"github.com/MichaelVenturi/go-practice-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// auth code
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		// abort:  stop instead of moving onto the next request handler
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
