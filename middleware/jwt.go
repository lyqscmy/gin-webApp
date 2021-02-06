// middleware/jwt.go

package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	//JWT Middleware for Gin framework
	jwt "github.com/appleboy/gin-jwt/v2"

	"api/models"
)

func JwtMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         "Realmname",
		Key:           []byte("SecrteKey"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: authenticator,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(200, gin.H{
				"code":    code,
				"message": message,
			})
		},
	}
}

//authenticator function for jwt middleware
func authenticator(c *gin.Context) (interface{}, error) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	//admin = password
	if user.Username == "admin" && user.Password == "password" {
		return user, nil
	}
	return nil, jwt.ErrFailedAuthentication
}
