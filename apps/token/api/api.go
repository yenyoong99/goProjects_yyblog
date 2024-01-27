package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/conf"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"github.com/yenyoong99/goProjects_yyblog/response"
)

func init() {
	ioc.Api().Registry(token.AppName, &TokenApiHandler{})
}

func NewTokenApiHandler(svc token.Service) *TokenApiHandler {
	return &TokenApiHandler{
		svc: svc,
	}
}

// TokenApiHandler RESTful interface
type TokenApiHandler struct {
	svc token.Service
}

func (h *TokenApiHandler) Init() error {
	h.svc = ioc.Controller().Get(token.AppName).(token.Service)
	return nil
}

func (h *TokenApiHandler) Destroy() error {
	return nil
}

// Registry
// Handle add path to http Server
// need root router, path prefix: /yyblog/api/v1
func (h *TokenApiHandler) Registry(rr gin.IRouter) {
	// all business module, need go Gin Engine to register the router
	// r := gin.Default()
	// rr := r.Group("/yyblog/api/v1")

	// module path
	// /yyblog/api/v1/tokens
	mr := rr.Group(token.AppName)
	mr.POST("/", h.Login)
	mr.DELETE("/", h.Logout)
}

func (h *TokenApiHandler) Login(c *gin.Context) {
	// 1. parse user request
	// http request can put at body, bytes
	// io.ReadAll(c.Request.Body)
	// defer c.Request.Body.Close()
	// json unmarshal json.Unmaral(body, o)

	// Body must be JSON type
	req := token.NewIssueTokenRequest("", "")
	if err := c.BindJSON(req); err != nil {
		response.Failed(c, err)
		return
	}

	// 2. logic processing
	tk, err := h.svc.IssueToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 2.1 set cookie
	c.SetCookie(
		token.TOKEN_COOKIE_KEY,
		tk.AccessToken,
		tk.AccessTokenExpiredAt,
		"/",
		conf.C().Application.Domain,
		false,
		true,
	)

	// 3. return process result
	response.Success(c, tk)
}

func (h *TokenApiHandler) Logout(c *gin.Context) {
	// For security reasons, the token is stored in the custom header obtained by Cookie.
	accessToken := token.GetAccessTokenFromHttp(c.Request)
	req := token.NewRevokeTokenRequest(accessToken, c.Query("refresh_token"))
	// 2. logic processing
	_, err := h.svc.RevokeToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 2.1 remove front-end cookie
	c.SetCookie(
		token.TOKEN_COOKIE_KEY,
		"",
		-1,
		"/",
		conf.C().Application.Domain,
		false,
		true,
	)

	// 3. return process result
	response.Success(c, "Logout Success")
}
