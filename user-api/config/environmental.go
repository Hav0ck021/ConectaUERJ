package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port                  = 0
	MysqlConnectionString = ""
	SecretKey             []byte
	FrontendURL           = ""
	EmailSender           = ""
	SMTPPort              = 0
	SMTPServer            = ""
	EMailSenderPassword   = ""
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	MysqlConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	FrontendURL = os.Getenv("FRONT_END_URL")

	SMTPPort, err = strconv.Atoi(os.Getenv("PORT_MAIL"))
	if err != nil {
		panic(".env not found")
	}

	SMTPServer = os.Getenv("SMTP_SERVER")
	EmailSender = os.Getenv("EMAIL_SENDER")
	EMailSenderPassword = os.Getenv("EMAIL_SENDER_PASSWORD")
}
