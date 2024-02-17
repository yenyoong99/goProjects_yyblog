package impl_test

import (
	"context"
	"testing"

	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/apps/token/impl"
	ui "github.com/yenyoong99/goProjects_yyblog/apps/user/impl"
)

var (
	i   token.Service
	ctx = context.Background()
)

/*
	{
	          "user_id": "9",
	          "username": "admin",
	          "access_token": "cmh62ncbajf1m8ddlpa0",
	          "access_token_expired_at": 7200,
	          "refresh_token": "cmh62ncbajf1m8ddlpag",
	          "refresh_token_expired_at": 28800,
	          "created_at": 1705140573,
	          "updated_at": 1705140573,
	          "role": 1
	}
*/
func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest("admin", "123456")
	req.RemindMe = true
	tk, err := i.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestRevokeToken(t *testing.T) {
	req := token.NewRevokeTokenRequest(
		"cmh62ncbajf1m8ddlpa0",
		"cmh62ncbajf1m8ddlpag",
	)
	tk, err := i.RevokeToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

// refresh token expired 8666.516636 minutes
/*
{
	"user_id": "9",
	"username": "admin",
	"access_token": "cmh63mkbajf1o5uh5cb0",
	"access_token_expired_at": 604800,
	"refresh_token": "cmh63mkbajf1o5uh5cbg",
	"refresh_token_expired_at": 604800,
	"created_at": 1705140698,
	"updated_at": 1705140698,
	"role": 0
}
*/
func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cmh63mkbajf1o5uh5cb0")
	tk, err := i.ValidateTokenRequest(ctx, req)
	// exception.IsException(err, token.ErrAccessTokenExpired)
	// if e, ok := err.(*exception.APIException); ok {
	// 	t.Log(e.String())
	// 	// TokenExpired Exceptions
	// 	if e.Code == token.ErrAccessTokenExpired.Code {
	// 		t.Log(e.String())
	// 	}
	// }

	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	i = impl.NewTokenServiceImpl(ui.NewUserServiceImpl())
}
