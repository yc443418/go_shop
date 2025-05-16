package users

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go_shop/database"
	"strings"
)

type Users struct {
	User_id      int    `gorm:"primarykey"`
	Username     string `gorm:"size:30"`
	Password     string `gorm:"size:128"`
	Email        string `gorm:"size:50"`
	Phone_number string `gorm:"size:11"`
}

func CheckLogin(username string, password string) (Users, error) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if username == "" {
		return Users{}, errors.New("用户名不能为空！")
	}

	if password == "" {
		return Users{}, errors.New("密码不能为空！")
	}

	var user Users

	database.Gdb.Where("username = ?", username).Find(&user)

	if user.Username == "" {
		return Users{}, errors.New("用户名错误！")
	}

	//md5加密
	hash := md5.Sum([]byte(password))
	if user.Password != hex.EncodeToString(hash[:]) {
		return Users{}, errors.New("密码错误！")
	}

	return user, nil
}

// 接收用户注册数据
type RegInfo struct {
	Username     string `form:"username" binding:"required,min=6,max=20"`
	Password     string `form:"password" binding:"required,min=6,max=20"`
	Password2    string `form:"password2" binding:"eqfield=Password"`
	Phone_number string `form:"phone" binding:"len=11"`
	Email        string `form:"email" binding:"email"`
	Captcha      string `form:"captcha" binding:"len=5"`
}

// 合理性验证用户注册信息
func CheckRegisterInfo(regInfo RegInfo) error {
	var user Users

	database.Gdb.Where("username = ?", regInfo.Username).First(&user)
	if user.User_id != 0 {
		return errors.New("该用户名已存在")
	}

	database.Gdb.Where("email = ?", regInfo.Email).First(&user)
	if user.User_id != 0 {
		return errors.New("该邮箱已注册过，可以尝试找回密码")
	}

	database.Gdb.Where("phone_number = ?", regInfo.Phone_number).First(&user)
	if user.User_id != 0 {
		return errors.New("该手机号已注册过，可以尝试找回密码")
	}

	user.Email = regInfo.Email
	user.Username = regInfo.Username
	user.Phone_number = regInfo.Phone_number
	//加密密码
	hash := md5.Sum([]byte(regInfo.Password))
	user.Password = hex.EncodeToString(hash[:])

	database.Gdb.Save(&user)

	return nil
}
