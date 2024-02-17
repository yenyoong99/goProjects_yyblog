package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"github.com/yenyoong99/goProjects_yyblog/response"
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
