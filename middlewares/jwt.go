package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/common/logging"
)

var jwtSecret = []byte("cloudfeet-jwt-token")
var logger = logging.GetLogger()

// Claims ...
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken ...
func GenerateToken(username string, password string) (string, error) {
	now := time.Now()
	// expireTime := now.Add(time.Hour * time.Duration(settings.Config.Jwt.ExpireHour))
	expireTime := now.Add(time.Hour * 3)

	claims := &Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "cloudfeet",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	logger.Info("gen token ", token)

	return token, err
}

// ParseToken ...
func ParseToken(token string) (*Claims, error) {
	// TODO: add db model query validate
	tokenClaims, err := jwt.ParseWithClaims(
		token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			logger.Info("parsed token = ", claims)
			return claims, nil
		}
	}

	return nil, err
}

// JwtMiddleware ...
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" && strings.Index(c.Request.URL.Path, "auth") != -1 {
			c.Next()
			return
		}

		token := c.Request.Header.Get("Token")
		logger.Info("get req token = ", token)
		_, err := ParseToken(token)
		if err != nil {
			logger.Debug(err.Error())
			c.JSON(http.StatusUnauthorized,
				gin.H{"code": 400, "msg": "auth failed with token or token expired", "data": nil})
			c.Abort()
			return
		}

		c.Next()
	}
}
