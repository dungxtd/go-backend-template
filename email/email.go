// Package email provides email sending functionality.
package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-mail/mail"
	"github.com/jaytaylor/html2text"
	"github.com/spf13/viper"
	"github.com/vanng822/go-premailer/premailer"
)

var (
	debug     bool
	templates *template.Template
)

type MailClient interface {
	LoginToken(name, address string, content ContentLoginToken) error
}

// Mailer is a SMTP mailer.
type Mailer struct {
	client *mail.Dialer
	from   Email
}

// NewMailer returns a configured SMTP Mailer.
func NewMailer(emailSMTPHost string, emailSMTPPort int, emailSMTPUser string, emailSMTPPassword string) (MailClient, error) {
	if err := parseTemplates(); err != nil {
		return nil, err
	}

	smtp := struct {
		Host     string
		Port     int
		User     string
		Password string
	}{
		emailSMTPHost,
		emailSMTPPort,
		emailSMTPUser,
		emailSMTPPassword,
	}

	s := &Mailer{
		client: mail.NewDialer(smtp.Host, smtp.Port, smtp.User, smtp.Password),
		from:   NewEmail(viper.GetString("email_from_name"), viper.GetString("email_from_address")),
	}

	if smtp.Host == "" {
		log.Println("SMTP host not set => printing emails to stdout")
		debug = true
		return s, nil
	}

	d, err := s.client.Dial()
	if err == nil {
		d.Close()
		return s, nil
	}
	return nil, err
}

// Send sends the mail via smtp.
func (m *Mailer) Send(email *message) error {
	if debug {
		log.Println("To:", email.to.Address)
		log.Println("Subject:", email.subject)
		log.Println(email.text)
		return nil
	}

	msg := mail.NewMessage()
	msg.SetAddressHeader("From", email.from.Address, email.from.Name)
	msg.SetAddressHeader("To", email.to.Address, email.to.Name)
	msg.SetHeader("Subject", email.subject)
	msg.SetBody("text/plain", email.text)
	msg.AddAlternative("text/html", email.html)

	return m.client.DialAndSend(msg)
}

// message struct holds all parts of a specific email message.
type message struct {
	from     Email
	to       Email
	subject  string
	template string
	content  interface{}
	html     string
	text     string
}

// parse parses the corrsponding template and content
func (m *message) parse() error {
	buf := new(bytes.Buffer)
	if err := templates.ExecuteTemplate(buf, m.template, m.content); err != nil {
		return err
	}
	prem, err := premailer.NewPremailerFromString(buf.String(), premailer.NewOptions())
	if err != nil {
		return err
	}

	html, err := prem.Transform()
	if err != nil {
		return err
	}
	m.html = html

	text, err := html2text.FromString(html, html2text.Options{PrettyTables: true})
	if err != nil {
		return err
	}
	m.text = text
	return nil
}

// Email struct holds email address and recipient name.
type Email struct {
	Name    string
	Address string
}

// NewEmail returns an email address.
func NewEmail(name string, address string) Email {
	return Email{
		Name:    name,
		Address: address,
	}
}

func parseTemplates() error {
	templates = template.New("").Funcs(fMap)
	return filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templates.ParseFiles(path)
			return err
		}
		return err
	})
}

var fMap = template.FuncMap{
	"formatAsDate":     formatAsDate,
	"formatAsDuration": formatAsDuration,
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d.%d.%d", day, month, year)
}

func formatAsDuration(t time.Time) string {
	dur := time.Until(t)
	hours := int(dur.Hours())
	mins := int(dur.Minutes())

	v := ""
	if hours != 0 {
		v += strconv.Itoa(hours) + " hours and "
	}
	v += strconv.Itoa(mins) + " minutes"
	return v
}
