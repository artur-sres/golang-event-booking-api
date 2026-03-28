package middlewares

import (
	"net/http"

	"github.com/artur-sres/golang-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func Authanticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
