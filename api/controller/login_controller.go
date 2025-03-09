package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/internal/resutil"
)

type LoginController struct {
	LoginUsecase      domain.LoginUsecase
	SocialAuthUsecase domain.SocialAuthUsecase
	Env               *bootstrap.Env
}

func (lc *LoginController) CheckUserExists(c *gin.Context) {
	var request domain.UserCheckRequest

	err := c.ShouldBind(&request)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if request.Email == "" && request.PhoneNumber == "" {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, nil, "Either email or phone number must be provided")
		return
	}

	var exists bool
	if request.Email != "" {
		_, err = lc.LoginUsecase.GetUserByEmail(c, request.Email)
		exists = err == nil
	} else {
		_, err = lc.LoginUsecase.GetUserByPhone(c, request.PhoneNumber)
		exists = err == nil
	}

	resutil.HandleDataResponse(c, http.StatusOK, domain.UserCheckResponse{Exists: exists})
}

func (lc *LoginController) generateTokens(user *domain.User) (string, string, error) {
	accessToken, err := lc.LoginUsecase.CreateAccessToken(user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (lc *LoginController) LoginWithEmail(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusNotFound, err, "User not found with the given email")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		resutil.HandleErrorResponse(c, http.StatusUnauthorized, nil, "Invalid credentials")
		return
	}

	accessToken, refreshToken, err := lc.generateTokens(&user)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	resutil.HandleDataResponse(c, http.StatusOK, loginResponse)
}

func (lc *LoginController) LoginWithPhone(c *gin.Context) {
	var request domain.LoginPhoneRequest

	err := c.ShouldBind(&request)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := lc.LoginUsecase.GetUserByPhone(c, request.PhoneNumber)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusNotFound, err, "User not found with the given phone number")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		resutil.HandleErrorResponse(c, http.StatusUnauthorized, nil, "Invalid credentials")
		return
	}

	accessToken, refreshToken, err := lc.generateTokens(&user)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	resutil.HandleDataResponse(c, http.StatusOK, loginResponse)
}

func (sc *LoginController) SocialLogin(c *gin.Context) {
	var request domain.SocialAuthRequest

	err := c.ShouldBind(&request)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var user *domain.User
	switch request.Provider {
	case "google":
		user, err = sc.SocialAuthUsecase.AuthenticateGoogle(c, request.AccessToken)
	case "facebook":
		user, err = sc.SocialAuthUsecase.AuthenticateFacebook(c, request.AccessToken)
	case "apple":
		user, err = sc.SocialAuthUsecase.AuthenticateApple(c, request.AccessToken)
	default:
		resutil.HandleErrorResponse(c, http.StatusBadRequest, nil, "Invalid provider")
		return
	}

	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	accessToken, refreshToken, err := sc.generateTokens(user)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	socialAuthResponse := domain.SocialAuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	resutil.HandleDataResponse(c, http.StatusOK, socialAuthResponse)
}
