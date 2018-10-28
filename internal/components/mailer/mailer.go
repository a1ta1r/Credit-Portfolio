package mailer

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func SendMail(reciever string, username string, password string) bool {
	mail := Mail{}
	mail.senderId = "credit-portfolio@mail.ru"
	mail.toIds = []string{reciever}
	mail.subject = "Регистрация в сервисе \"Кредитный Портфель\""
	mail.body = "Данное письмо отправлено вам, так как оно указано при регистрации в сервисе \"Кредитный портфель\".\nВаше имя пользователя для входа:\n        " + username + "\nВаш пароль для входа:\n        " + password + "\n\nДобро пожаловать!"

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.mail.ru", port: "465"}

	log.Println(smtpServer.host)

	auth := smtp.PlainAuth("", mail.senderId, "Pipi1597", smtpServer.host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Println(err)
		return false
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Println(err)
		return false
	}

	if err = client.Auth(auth); err != nil {
		log.Println(err)
		return false
	}

	if err = client.Mail(mail.senderId); err != nil {
		log.Println(err)
		return false
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Println("Such email does not exist!")
			log.Println(err)
			return false
		}
	}

	w, err := client.Data()
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Println(err)
		return false
	}

	err = w.Close()
	if err != nil {
		log.Println(err)
		return false
	}

	client.Quit()

	log.Println("Mail sent successfully")
	return true

}
