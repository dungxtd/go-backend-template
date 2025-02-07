package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`

	//Auth
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`

	//SMTP email
	EmailSMTPHost          string `mapstructure:"EMAIL_SMTP_HOST"`
	EmailSMTPPort          int    `mapstructure:"EMAIL_SMTP_PORT"`
	EmailSMTPUser          string `mapstructure:"EMAIL_SMTP_USER"`
	EmailSMTPPassword      string `mapstructure:"EMAIL_SMTP_PASSWORD"`
	EmailFromName          string `mapstructure:"EMAIL_FROM_NAME"`
	EmailFromAddress       string `mapstructure:"EMAIL_FROM_ADDRESS"`
	SignupMailExpiryMinute int    `mapstructure:"SIGNUP_MAIL_EXPIRY_MINUTE"`

	//Twilio
	TwilioAccountSID  string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken   string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioPhoneNumber string `mapstructure:"TWILIO_PHONE_NUMBER"`

	//Unimtx
	UnimtxAccessKeyID     string `mapstructure:"UNIMTX_ACCESS_KEY_ID"`
	UnimtxAccessKeySecret string `mapstructure:"UNIMTX_ACCESS_KEY_SECRET"`

	//SpeedSMS
	SpeedSmsToken string `mapstructure:"SPEEDSMS_TOKEN"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
