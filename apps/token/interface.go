package token

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	AppName = "tokens"
)

// Service Token Service
type Service interface {
	// IssueToken login: issue token
	IssueToken(ctx context.Context, request *IssueTokenRequest) (*Token, error)

	// RevokeToken logout: revoke token
	RevokeToken(ctx context.Context, request *RevokeTokenRequest) (*Token, error)

	// ValidateTokenRequest verify: token
	ValidateTokenRequest(ctx context.Context, request *ValidateTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
	}

}

// IssueTokenRequest issue token
type IssueTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// extend token expired time to 1 week
	RemindMe bool `json:"remind_me"`
}

func NewRevokeTokenRequest(accessToken, refreshToken string) *RevokeTokenRequest {
	return &RevokeTokenRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

}

// RevokeTokenRequest revoke token
type RevokeTokenRequest struct {
	// issue the AccessToken to user (need token to access)
	AccessToken string
	// revoke the token need to make sure is paired with RefreshToken
	RefreshToken string
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: accessToken,
	}

}

type ValidateTokenRequest struct {
	AccessToken string
}

func GetAccessTokenFromHttp(req *http.Request) string {
	// custom header, Ex:Authorization: xxx Bearer xxx
	ah := req.Header.Get(TOKEN_HEADER_KEY)
	if ah != "" {
		hv := strings.Split(ah, " ")
		if len(hv) > 1 {
			return hv[1]
		}
	}

	// Cookies
	ck, err := req.Cookie(TOKEN_COOKIE_KEY)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	val, _ := url.QueryUnescape(ck.Value)
	return val

}
