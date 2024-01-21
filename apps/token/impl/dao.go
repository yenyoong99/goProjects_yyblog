package impl

import (
	"context"

	"goProjects/yyblog/apps/token"
)

// Data operations that need to be reused
func (i *TokenServiceImpl) getToken(ctx context.Context, accessToken string) (*token.Token, error) {
	tk := token.NewToken(false)
	err := i.db.
		WithContext(ctx).
		Model(&token.Token{}).
		Where("access_token = ?", accessToken).
		First(tk).
		Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}
