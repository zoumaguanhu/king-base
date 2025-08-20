package secret

import (
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/constants"
	"os"
	"path/filepath"
)

func Parse(mode string) (string, bool) {
	logx.Infof("current mode:%v", mode)
	if mode != constants.PRO_MODE {
		return constants.EMPTY_STRING, false
	}
	// 获取环境变量 DB_PASSWORD_FILE，默认使用 /run/secrets/db-password
	passwordFile := os.Getenv(constants.DB_PASSWORD_FILE)

	// 确保路径是绝对路径（可选，视需求而定）
	absPath, err := filepath.Abs(passwordFile)
	if err != nil {
		logx.Infof("passwordFile filepath err:%v", err)
		return constants.EMPTY_STRING, false
	}

	// 读取文件内容
	data, err := os.ReadFile(absPath)
	if err != nil {
		logx.Infof("passwordFile ReadFile err:%v", err)
		return constants.EMPTY_STRING, false
	}
	// 去除可能的换行符或空格
	return string(data), true
}
