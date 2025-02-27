package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	MongoDBHost    string `mapstructure:"MONGO_DB_HOST"`
	MongoDBPort    string `mapstructure:"MONGO_DB_PORT"`
	MongoDBUser    string `mapstructure:"MONGO_DB_USER"`
	MongoDBPass    string `mapstructure:"MONGO_DB_PASS"`
	MongoDBName    string `mapstructure:"MONGO_DB_NAME"`

	PostgresDBHost  string `mapstructure:"POSTGRES_DB_HOST"`
	PostgresDBPort  string `mapstructure:"POSTGRES_DB_PORT"`
	PostgresDBUser  string `mapstructure:"POSTGRES_DB_USER"`
	PostgresDBPass  string `mapstructure:"POSTGRES_DB_PASS"`
	PostgresDBName  string `mapstructure:"POSTGRES_DB_NAME"`
	PostgresSSLMode string `mapstructure:"POSTGRES_SSL_MODE"`

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

	//Minio
	MinioEndpoint    string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKeyID string `mapstructure:"MINIO_ACCESS_KEY_ID"`
	MinioSecretKey   string `mapstructure:"MINIO_SECRET"`
	MinioUseSSL      bool   `mapstructure:"MINIO_USE_SSL"`
}

func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}
}

func NewEnv() *Env {
	env := Env{}

	err := viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
