package flags

import (
	"fmt"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/lunarise-dev/lunar-gate/model"
	"github.com/lunarise-dev/lunar-gate/utils/pwd"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type User struct {
}

func (User) Create() {
	fmt.Println("请输入用户名")
	var username string
	fmt.Scanln(&username)
	var user model.User
	err := global.DB.Take(&user, "username = ?", username).Error
	if err == nil {
		logrus.Errorf("用户名已存在")
		return
	}

	fmt.Println("请输入密码")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		logrus.Errorf("密码读取失败 %s", err)
		return
	}
	fmt.Println("请再次输入密码")
	rePassword, err := terminal.ReadPassword(int(os.Stdin.Fd())) // 读取用户输入的密码
	if err != nil {
		logrus.Errorf("密码读取失败 %s", err)
		return
	}
	if string(password) != string(rePassword) {
		logrus.Errorf("两次密码不一致")
		return
	}

	hashPwd := pwd.Hash(string(password))
	err = global.DB.Create(&model.User{
		Username: username,
		Password: hashPwd,
		IsAdmin:  true,
	}).Error
	if err != nil {
		logrus.Errorf("用户创建失败 %s", err)
		return
	}
	logrus.Infof("用户创建成功")
}
