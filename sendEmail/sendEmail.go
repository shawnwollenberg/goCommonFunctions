package SendEmail

import (
	"bytes"
	"net/smtp"
)

func SendEmail(to, from, fromName, subject, msg string) {
	c, err := smtp.Dial("mail.mt.com:25")
	if err != nil {
		panic(err)
	}

	if err := c.Mail(from); err != nil {
		panic(err)
	}
	if err := c.Rcpt(to); err != nil {
		panic(err)
	}

	wc, err := c.Data()
	if err != nil {
		panic(err)
	}
	defer wc.Close()
	defer c.Quit()
	buf := bytes.NewBufferString("To: " + to + "\r\n" +
		"From: " + fromName + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		msg + "\r\n")
	if _, err = buf.WriteTo(wc); err != nil {
		panic(err)
	}

}
