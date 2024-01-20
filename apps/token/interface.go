package token

import (
	"context"
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
	Username string
	Password string
	// extend token expired time to 1 week
	RemindMe bool
}

// RevokeTokenRequest revoke token
type RevokeTokenRequest struct {
	// issue the AccessToken to user (need token to access)
	AccessToken string
	// revoke the token need to make sure is paired with RefreshToken
	RefreshToken string
}

type ValidateTokenRequest struct {
	AccessToken string
}
