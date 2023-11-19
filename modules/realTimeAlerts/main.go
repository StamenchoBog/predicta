package realTimeAlerts

import (
	"fmt"
	"log"
	"net/smtp"
)

// Constants
const alertSubject = "Security Alert: Anomaly Detected"
const alertBody = "Anomaly detected in the network. Please investigate."

func SendEmail(from string, password string, to string, smtpServer string, smtpPort int, subject, body string) {
	auth := smtp.PlainAuth("", from, password, smtpServer)
	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
