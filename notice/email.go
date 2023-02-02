/**
 * Created by GoLand
 * @file   email.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/17 00:03
 * @desc   email.go
 */

package notice

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

// EmailUtil 邮件实体
type EmailUtil struct {
	from     string
	to       []string
	cc       []string
	bcc      []string
	subject  string
	content  string
	addr     string
	userName string
	password string
	host     string
}

// Init 初始化
func (m *EmailUtil) Init(from string, userName string, password string, addr string, host string) *EmailUtil {
	m.userName = userName
	m.from = from
	m.addr = addr
	m.password = password
	m.host = host
	return m
}

// SetTo 设置接收方的邮箱
func (m *EmailUtil) SetTo(to []string) *EmailUtil {
	m.to = to
	return m
}

// SetCc 设置抄送如果抄送多人逗号隔开
func (m *EmailUtil) SetCc(cc []string) *EmailUtil {
	m.cc = cc
	return m
}

// SetBcc 设置秘密抄送
func (m *EmailUtil) SetBcc(bcc []string) *EmailUtil {
	m.bcc = bcc
	return m
}

// SetSubject 设置主题
func (m *EmailUtil) SetSubject(subject string) *EmailUtil {
	m.subject = subject
	return m
}

// SetContent 设置文件发送的内容
func (m *EmailUtil) SetContent(content string) *EmailUtil {
	m.content = content
	return m
}

// SendEmail 发送邮件
func (m *EmailUtil) SendEmail() (bool, error) {
	e := email.NewEmail()
	// 设置发送方的邮箱
	e.From = m.from
	// 设置接收方的邮箱
	e.To = m.to
	// 设置抄送如果抄送多人逗号隔开
	e.Cc = m.cc
	// 设置秘密抄送
	e.Bcc = m.bcc
	// 设置主题
	e.Subject = m.subject
	// 设置文件发送的内容
	e.Text = []byte(m.content)
	// 设置服务器相关的配置
	// err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "244395692@qq.com", "rtrohqzojhtrcajf", "smtp.qq.com"))
	err := e.Send(m.addr, smtp.PlainAuth("", m.userName, m.password, m.host))
	if err != nil {
		return false, err
	}
	return true, nil
}
