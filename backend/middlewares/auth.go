package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/utils"
	"github.com/vallezw/RomManager/backend/utils/token"
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
