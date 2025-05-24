package mail

import (
	"fmt"
	"king.com/king/base/common/strs"
	"testing"
)

func TestSysMail(t *testing.T) {
	m := &MailConf{
		Host:     "smtp.126.com",
		Port:     465,
		Username: "qydkkww@126.com",
		Password: "SYufBPBXZAUXQxaY",
		SSL:      true,
	}
	s := strs.ObjToStr(m)
	fmt.Printf("TestSysMail result:%v", s)
}
