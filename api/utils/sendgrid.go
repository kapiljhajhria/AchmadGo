package utils

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/samhj/AchmadGo/api/config"
	"github.com/samhj/AchmadGo/api/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//SendEmail ...
func SendEmail(email *models.Email) error{

	from := mail.NewEmail("Zouq", email.SenderEmail)
	subject := email.Subject
	to := mail.NewEmail(email.ReceiverName, email.ReceiverEmail)
	header,_ := ReadFile("header")
	footer,_ := ReadFile("footer")
	content, _ := ReadFile(email.FileName)
	
	htmlContent := header + Replacer(content, email.Replacer) + footer
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(config.Config("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println("Error Sending Email:", err)
	} else {
		fmt.Println("Email StatusCode:", response.StatusCode)
	}

	return err

}

//Replacer replaces the contents in the email template
func Replacer(s string, rep map[string]string) string {
	content := s
	for k, v := range rep {
		content = strings.ReplaceAll(content, k, v)
	}
	return content
}

//ReadFile ...
func ReadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile("./public/emailTemplates/" + filename + ".html") // just pass the file name
	if err != nil {
		fmt.Println("Reading Email Template Error For "+filename+":", err.Error())
	}
	return string(b), err                                                      // convert content to a 'string'
}
