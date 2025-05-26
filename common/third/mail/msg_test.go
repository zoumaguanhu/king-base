package mail

import (
	"fmt"
	"king.com/king/base/common/strs"
	"testing"
)

func TestSysMail(t *testing.T) {
	m := &MailConf{
		Host:           "smtp.126.com",
		Port:           465,
		Username:       "qydkkww@126.com",
		Password:       "SYufBPBXZAUXQxaY",
		SSL:            true,
		WorkerNum:      2,
		JobQueue:       100,
		WaitTimeOut:    5,
		RetryCount:     1,
		MaxCountPerDay: 1000,
	}
	s := strs.ObjToStr(m)
	fmt.Printf("TestSysMail result:%v", s)
}
