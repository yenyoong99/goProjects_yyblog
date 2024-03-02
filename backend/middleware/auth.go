package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"github.com/yenyoong99/goProjects_yyblog/response"
	"time"
)

func Auth(c *gin.Context) {
	svc := ioc.Controller().Get(token.AppName).(token.Service)

	tk := token.GetAccessTokenFromHttp(c.Request)
	req := token.NewValidateTokenRequest(tk)
	tkObj, err := svc.ValidateTokenRequest(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	c.Set(token.TOKEN_MIDDLEWARE_KEY, tkObj)

	// c.Get(token.TOKEN_MIDDLEWARE_KEY).(*token.Token).UserName
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
