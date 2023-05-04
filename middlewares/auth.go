package middlewares

import (
	"errors"
	"net/http"

	"github.com/RomManager/server/utils"
	"github.com/RomManager/server/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			utils.DoError(c, http.StatusUnauthorized, errors.New("Unauthorized"))
			c.Abort()
			return
		}
		c.Next()
	}
}
