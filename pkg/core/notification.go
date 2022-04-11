package core

import (
	"errors"
	"fmt"

	"gopkg.in/gomail.v2"
)

type Email struct {
	dialer *gomail.Dialer
}

func newEmail() *Email {
	if CFG.Notification.Smtp == nil {
		return nil
	}
	e := Email{}
	e.dialer = gomail.NewDialer(CFG.Notification.Smtp.Host, CFG.Notification.Smtp.Port,
		CFG.Notification.Smtp.Username, CFG.Notification.Smtp.Password)
	return &e
}

func (e *Email) Send(receiver, title, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", CFG.Notification.Smtp.Username)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", content)
	return e.dialer.DialAndSend(m)
}

var EMAIL *Email

func init() {
	STARTUP_HOOKS = append(STARTUP_HOOKS, init_)
}
func init_() {
	EMAIL = newEmail()
}

// 发送激活链接信息
func SendActivation(email, activationToken string) error {
	if EMAIL == nil {
		return errors.New("notification.smtp not found")
	}
	return EMAIL.Send(email, "你的帐号正待激活！", fmt.Sprintf(`
	<h1>勇敢的尝试是成功的一半 <br />A bold attempt is half success</h1>
	<a href="https://%s/activation/%s">https://%s/activation/%s</a>
	`, CFG.Site.Domain, activationToken, CFG.Site.Domain, activationToken))
}

// 发送重置密码链接
func SendResetPassword(email, activationToken string) error {
	if EMAIL == nil {
		return errors.New("notification.smtp not found")
	}
	return EMAIL.Send(email, "重置密码！", fmt.Sprintf(`
	<h1>永远不要、不要、不要、不要放弃<br />Never, never, never, never give up</h1>
	<a href="https://%s/reset-password/%s">https://%s/reset-password/%s</a>
	`, CFG.Site.Domain, activationToken, CFG.Site.Domain, activationToken))
}
