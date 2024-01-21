package token

import (
	"fmt"
	"github.com/infraboard/mcube/tools/pretty"
	"time"

	"github.com/rs/xid"
	"github.com/yenyoong99/goProjects_yyblog/apps/user"
)

const (
	DefaultExpiredAt = 2 * 60 * 60
	WeekExpiredAt    = 7 * 24 * 60 * 60
)

func NewToken(remindMe bool) *Token {
	accessTokenExpiredTime := DefaultExpiredAt

	if remindMe {
		accessTokenExpiredTime = WeekExpiredAt
	}

	return &Token{
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  accessTokenExpiredTime,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: accessTokenExpiredTime * 4,
		CreatedAt:             time.Now().Unix(),
	}
}

type Token struct {
	UserId                string    `json:"user_id" gorm:"column:user_id"`
	UserName              string    `json:"username" gorm:"column:username"`
	AccessToken           string    `json:"access_token" gorm:"column:access_token"`
	AccessTokenExpiredAt  int       `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`
	RefreshToken          string    `json:"refresh_token" gorm:"column:refresh_token"`
	RefreshTokenExpiredAt int       `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`
	CreatedAt             int64     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt             int64     `json:"updated_at" gorm:"column:updated_at"`
	Role                  user.Role `json:"role" gorm:"-"`
}

func (t *Token) CheckRefreshToken(refreshToken string) error {
	if t.RefreshToken != refreshToken {
		return fmt.Errorf("refresh token incorrect")
	}
	return nil
}

func (t *Token) ValidateExpired() error {
	refreshExpiredTime := time.
		Unix(t.CreatedAt, 0).
		Add(time.Duration(t.RefreshTokenExpiredAt) * time.Second)

	rDelta := time.Since(refreshExpiredTime).Minutes()
	if rDelta > 0 {
		return ErrRefreshTokenExpired.WithMessagef("refresh token expired %f minutes", rDelta)
	}

	accessExpiredTime := time.
		Unix(t.CreatedAt, 0).
		Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)

	aDelta := time.Since(accessExpiredTime).Minutes()
	if aDelta > 0 {
		return ErrAccessTokenExpired.WithMessagef("access token expired %f minutes", aDelta)
	}
	return nil
}

func (t *Token) TableName() string {
	return "tokens"
}

func (t *Token) String() string {
	return pretty.ToJSON(t)
}
