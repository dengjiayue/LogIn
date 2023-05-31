package funcmod

import (
	"fmt"

	"github.com/badoux/checkmail"
	"gopkg.in/gomail.v2"
)

// 使用gomail库创建SMTP客户端
func CreatDialer() *gomail.Dialer {
	return gomail.NewDialer("smtp.qq.com", 465, "你的邮箱", "你的授权码") //AuthCode为邮箱的授权码
}

// 发送邮件函数,+邮箱可用性验证
func SendEmail(dialer *gomail.Dialer, to string, verification int) error {
	// 发送邮件的QQ邮箱地址
	qqEmail := "你的邮箱"

	// 检查电子邮件地址是否可用
	err := checkmail.ValidateFormat(to)
	if err != nil {
		return fmt.Errorf("email address %s is not available: %s", to, err.Error())
	}
	// 创建邮件消息
	message := gomail.NewMessage()
	message.SetHeader("From", qqEmail)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "验证码")
	message.SetBody("text/plain", fmt.Sprintf("验证码:%d", verification))

	// 发送邮件消息,开携程发生邮件
	go dialer.DialAndSend(message)

	return nil
}
