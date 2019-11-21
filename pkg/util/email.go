package util

import (
	"fmt"
	"os"

	"github.com/gethinyan/sso-server/pkg/setting"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendEmail 发送验证码邮件方法
func SendEmail(email string, userName string, code int) error {
	from := mail.NewEmail(setting.Email.FromSubject, setting.Email.From)
	subject := setting.Email.Subject
	to := mail.NewEmail(userName, email)
	plainTextContent := setting.Email.Subject
	htmlContent := fmt.Sprintf("<p>尊敬的Sso-server用户%s：</p><p>您好!</p><p>您的验证码是：%d。</p>", userName, code)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	fmt.Println(os.Getenv("SENDGRID_API_KEY"))
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return nil
}
