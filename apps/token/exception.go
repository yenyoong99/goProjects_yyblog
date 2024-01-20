package token

import (
	"goProjects/yyblog/exception"
)

var (
	ErrAccessTokenExpired  = exception.NewAPIException(5000, "access token expired")
	ErrRefreshTokenExpired = exception.NewAPIException(5001, "refresh token expired")
)
