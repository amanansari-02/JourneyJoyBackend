package email

import (
	"JourneyJoyBackend/src/models"
	"os"
	"strconv"

	mail "gopkg.in/mail.v2"
)

func SendBookingConfirmationMessage(Booking models.Booking) error {
	emailHost := os.Getenv("EMAIL_HOST")
	emailPort := os.Getenv("EMAIL_PORT")
	emailUser := os.Getenv("EMAIL_USER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailTo := Booking.Email
	prop_type := Booking.Property.PropertyType
	priceStr := strconv.FormatInt(Booking.Price, 10)

	port, err := strconv.Atoi(emailPort)
	if err != nil {
		return nil
	}

	d := mail.NewDialer(emailHost, port, emailUser, emailPassword)

	m := mail.NewMessage()
	m.SetHeader("From", emailUser)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "Enquiry data from JourneyJoy")
	m.SetBody("text/plain", "Booking confirmed successfully:\n\nCustomer Name: "+
		Booking.FullName+"\nCustomer Email: "+emailTo+"\n "+prop_type+" name is "+Booking.Property.PropertyName+
		"\n Your "+prop_type+" booked "+Booking.StartDate+" to "+Booking.EndDate+
		"\n Total price: "+priceStr)

	return d.DialAndSend(m)
}
