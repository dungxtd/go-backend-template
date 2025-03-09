package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/internal/httputil"
)

type socialAuthUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	httpClient     *httputil.Client
	env            *bootstrap.Env
}

func NewSocialAuthUsecase(userRepository domain.UserRepository, timeout time.Duration, env *bootstrap.Env) domain.SocialAuthUsecase {
	return &socialAuthUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		httpClient:     httputil.NewClient(timeout),
		env:            env,
	}
}

func (sa *socialAuthUsecase) AuthenticateGoogle(c context.Context, token string) (*domain.User, error) {
	googleUserInfo, err := sa.getGoogleUserInfo(token)
	if err != nil {
		return nil, err
	}

	user, err := sa.userRepository.GetByEmail(c, googleUserInfo.Email)
	if err == nil {
		// User exists, update Google ID if not set
		if user.GoogleID == "" {
			user.GoogleID = googleUserInfo.ID
			// Update user in repository
			if err := sa.userRepository.Update(c, &user); err != nil {
				return nil, err
			}
		}
		return &user, nil
	}

	// Create new user
	newUser := &domain.User{
		ID:        uuid.New().String(),
		Name:      googleUserInfo.Name,
		Email:     googleUserInfo.Email,
		GoogleID:  googleUserInfo.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = sa.userRepository.Create(c, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (sa *socialAuthUsecase) AuthenticateFacebook(c context.Context, token string) (*domain.User, error) {
	fbUserInfo, err := sa.getFacebookUserInfo(token)
	if err != nil {
		return nil, err
	}

	user, err := sa.userRepository.GetByEmail(c, fbUserInfo.Email)
	if err == nil {
		// User exists, update Facebook ID if not set
		if user.FacebookID == "" {
			user.FacebookID = fbUserInfo.ID
			// Update user in repository
			if err := sa.userRepository.Update(c, &user); err != nil {
				return nil, err
			}
		}
		return &user, nil
	}

	// Create new user
	newUser := &domain.User{
		ID:         uuid.New().String(),
		Name:       fbUserInfo.Name,
		Email:      fbUserInfo.Email,
		FacebookID: fbUserInfo.ID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = sa.userRepository.Create(c, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (sa *socialAuthUsecase) AuthenticateApple(c context.Context, token string) (*domain.User, error) {
	appleUserInfo, err := sa.getAppleUserInfo(token)
	if err != nil {
		return nil, err
	}

	user, err := sa.userRepository.GetByEmail(c, appleUserInfo.Email)
	if err == nil {
		// User exists, update Apple ID if not set
		if user.AppleID == "" {
			user.AppleID = appleUserInfo.ID
			// Update user in repository
			if err := sa.userRepository.Update(c, &user); err != nil {
				return nil, err
			}
		}
		return &user, nil
	}

	// Create new user
	newUser := &domain.User{
		ID:        uuid.New().String(),
		Name:      appleUserInfo.Name,
		Email:     appleUserInfo.Email,
		AppleID:   appleUserInfo.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = sa.userRepository.Create(c, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Helper structs for social provider responses
type googleUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type facebookUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type appleUserInfo struct {
	ID    string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Helper functions to get user info from providers
func (sa *socialAuthUsecase) getGoogleUserInfo(token string) (*googleUserInfo, error) {
	var userInfo googleUserInfo
	err := sa.httpClient.Get(
		context.Background(),
		fmt.Sprintf("%s?access_token=%s", sa.env.GoogleUserInfoURL, token),
		&userInfo,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info from Google: %w", err)
	}
	return &userInfo, nil
}

func (sa *socialAuthUsecase) getFacebookUserInfo(token string) (*facebookUserInfo, error) {
	var userInfo facebookUserInfo
	err := sa.httpClient.Get(
		context.Background(),
		fmt.Sprintf("%s?fields=id,name,email&access_token=%s", sa.env.FacebookUserInfoURL, token),
		&userInfo,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info from Facebook: %w", err)
	}
	return &userInfo, nil
}

func (sa *socialAuthUsecase) getAppleUserInfo(token string) (*appleUserInfo, error) {
	// Implementation for Apple Sign In validation
	// This would involve validating the identity token and extracting user information
	// The actual implementation would require Apple's authentication libraries
	return nil, errors.New("Apple Sign In validation not implemented")
}
