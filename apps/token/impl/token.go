package impl

import (
	"context"
	"fmt"

	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/apps/user"
)

// IssueToken
// IssueToken Login: Issue AccessToken
// depend User module to validate (username and password)
func (i *TokenServiceImpl) IssueToken(
	ctx context.Context,
	in *token.IssueTokenRequest) (
	*token.Token, error) {
	// 1.1 confirm user password is correct
	req := user.NewQueryUserRequest()
	req.Username = in.Username
	// interface oriented, Do abstract programming
	us, err := i.user.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(us.Items) == 0 {
		return nil, fmt.Errorf("username or password incorrect")
	}

	// 1.2 verify the user password
	u := us.Items[0]
	if err := u.CheckPassword(in.Password); err != nil {
		return nil, err
	}

	// 2. if correct, issue token to user
	tk := token.NewToken(in.RemindMe)
	// Associated user information
	tk.UserId = fmt.Sprintf("%d", u.Id)
	tk.UserName = u.Username
	tk.Role = u.Role

	// 3. keep token
	err = i.db.WithContext(ctx).
		Model(&token.Token{}).
		Create(tk).
		Error
	if err != nil {
		return nil, err
	}

	return tk, nil
}

// RevokeToken
// RevokeToken Logout: revokeToken, delete the user token
func (i *TokenServiceImpl) RevokeToken(
	ctx context.Context,
	in *token.RevokeTokenRequest) (
	*token.Token, error) {
	// query user token
	tk, err := i.getToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	// Refresh confirm
	err = tk.CheckRefreshToken(in.RefreshToken)
	if err != nil {
		return nil, err
	}

	// query token
	// DELETE FROM `tokens` WHERE access_token = 'cmh62ncbajf1m8ddlpa0' AND refresh_token = 'cmh62ncbajf1m8ddlpag'
	err = i.db.WithContext(ctx).
		Where("access_token = ?", in.AccessToken).
		Where("refresh_token = ?", in.RefreshToken).
		Delete(&token.Token{}).
		Error
	if err != nil {
		return nil, err
	}

	return tk, nil
}

// ValidateTokenRequest
// ValidateToken verify user token
func (i *TokenServiceImpl) ValidateTokenRequest(
	ctx context.Context,
	in *token.ValidateTokenRequest) (
	*token.Token, error) {
	// 1. query token, check token is exists
	tk, err := i.getToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	// 2. check token expired
	if err := tk.ValidateExpired(); err != nil {
		return nil, err
	}

	// 3. token legal, return
	return tk, nil
}
