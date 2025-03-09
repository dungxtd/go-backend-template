package domain

import (
	"context"
)

type SocialAuthRequest struct {
	Provider    string `form:"provider" binding:"required,oneof=google facebook apple"`
	AccessToken string `form:"access_token" binding:"required"`
}

type SocialAuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SocialAuthUsecase interface {
	AuthenticateGoogle(c context.Context, token string) (*User, error)
	AuthenticateFacebook(c context.Context, token string) (*User, error)
	AuthenticateApple(c context.Context, token string) (*User, error)
}
