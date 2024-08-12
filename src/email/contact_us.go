package email

import (
	"JourneyJoyBackend/src/models"
	"os"
	"strconv"

	mail "gopkg.in/mail.v2"
)

func SendContactUsEmail(contactUs models.ContactUs) error {
	emailHost := os.Getenv("EMAIL_HOST")
	emailPort := os.Getenv("EMAIL_PORT")
	emailUser := os.Getenv("EMAIL_USER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailFrom := contactUs.Email

	port, err := strconv.Atoi(emailPort)
	if err != nil {
		return nil
	}

	d := mail.NewDialer(emailHost, port, emailUser, emailPassword)

	m := mail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("Reply-To", contactUs.Email)
	m.SetHeader("To", emailUser)
	m.SetHeader("Subject", "Enquiry data from JourneyJoy")
	m.SetBody("text/plain", "New contact us message received:\n\nName: "+contactUs.FullName+"\nEmail: "+contactUs.Email+"\nMessage: "+contactUs.Message)

	return d.DialAndSend(m)
}
