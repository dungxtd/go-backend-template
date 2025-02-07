package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/email"
	"github.com/sportgo-app/sportgo-go/sms"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
	Mailer        email.MailClient
	UnimtxClient  sms.UnimtxClient
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	go func() {
		sc.UnimtxClient.SendOTP("+84976447558")
	}()

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

	c.JSON(http.StatusOK, signupResponse)
}
