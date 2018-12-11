package core

import (
	"github.com/devectron/sunlight/log"
	. "github.com/mailjet/mailjet-apiv3-go"
)

func SendMail(tomail string, link string, publicapi string, privateapi string) {
	log.Inf("Sending e-mail to %s", tomail)
	mailjetClient := NewMailjetClient(publicapi, privateapi)
	email := &InfoSendMail{
		FromEmail: "devectron.not.replay@gmail.com",
		FromName:  "Devectron Team",
		Subject:   "File converted successfully",
		TextPart:  "",
		HTMLPart:  "<h1>Dear " + tomail + ":</h1> <p>Your file converted successfully you can download it from here <a href=\"" + link + "\">LINK</a> your file will be deleted after <b>5 min</b></p>",
		Recipients: []Recipient{
			Recipient{
				Email: tomail,
			},
		},
	}
	_, err := mailjetClient.SendMail(email)
	if err != nil {
		log.Err("Error While sending email %v", err)
	}
	log.Inf("Mail send successfully to %s", tomail)
}
