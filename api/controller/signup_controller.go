package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/email"
	"github.com/sportgo-app/sportgo-go/internal/resutil"
	"github.com/sportgo-app/sportgo-go/sms"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
	Mailer        email.MailClient
	SmsAdapter    sms.SmsAdapter
}

func (sc *SignupController) SignupWithEmail(c *gin.Context) {

	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		resutil.HandleErrorResponse(c, http.StatusConflict, nil, "User already exists with the given email")
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       uuid.New().String(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	//tokenURL, _ := url.JoinPath("", "lt.Token")

	//go func() {
	//	content := email.ContentLoginToken{
	//		Email:  request.Email,
	//		Name:   request.Name,
	//		URL:    tokenURL,
	//		Token:  "lt.Token",
	//		Expiry: time.Now().Add(time.Minute * time.Duration(sc.Env.SignupMailExpiryMinute)),
	//	}
	//	if err := sc.Mailer.LoginToken(request.Name, request.Email, content); err != nil {
	//		log.Fatal("Send email error: ", err)
	//	}
	//}()

	resutil.HandleDataResponse(c, http.StatusOK, signupResponse)
}

func (sc *SignupController) SignupWithPhone(c *gin.Context) {
	var request domain.SignupPhoneRequest

	err := c.ShouldBind(&request)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	_, err = sc.SignupUsecase.GetUserByPhone(c, request.PhoneNumber)
	if err == nil {
		resutil.HandleErrorResponse(c, http.StatusConflict, nil, "User already exists with the given phone number")
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	user := domain.User{
		ID:          uuid.New().String(),
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Password:    string(encryptedPassword),
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	resutil.HandleDataResponse(c, http.StatusOK, signupResponse)
}
