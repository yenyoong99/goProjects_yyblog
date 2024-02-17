package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/apps/user"
	"github.com/yenyoong99/goProjects_yyblog/exception"
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

// Required Define user role permission
func Required(roles ...user.Role) gin.HandlerFunc {
	return func(context *gin.Context) {
		// verify user role
		// token get user role, define the interface role
		if v, ok := context.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
			hasPerm := false
			// traverse to determine whether the user is in the role list
			for _, r := range roles {
				//fmt.Println(roles, v.(*token.Token).Role)
				if r == v.(*token.Token).Role {
					hasPerm = true
				}
			}

			if !hasPerm {
				response.Failed(context, exception.ErrPermissionDeny.WithMessagef("user role not allow access"))
				return
			}
		} else {
			response.Failed(context, exception.ErrUnauthorized)
		}
	}
}
