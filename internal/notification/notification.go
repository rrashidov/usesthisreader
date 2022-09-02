package notification

import "net/smtp"

type NotificationClient interface {
	Notify() error
}

type MailNotificationClient struct {
	Host     string
	Port     string
	Mail     string
	Password string
}

func (client MailNotificationClient) Notify() error {

	from := client.Mail
	password := client.Password

	toEmailAddress := client.Mail
	to := []string{toEmailAddress}

	host := client.Host //"smtp.gmail.com"
	port := client.Port //"587"
	address := host + ":" + port

	subject := "New article in usesthis.com"
	body := "There is a new article in usesthis.com. Go check it out!"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)

	return err
}
