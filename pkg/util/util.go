package util

import (
	"math/rand"
	"time"

	"github.com/gethinyan/sso/pkg/setting"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString 生成随机字符串
func RandomString(n int) string {
	str := make([]rune, n)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

// RandomCode 生成随机验证码
func RandomCode() int {
	min := setting.Code.Min
	max := setting.Code.Max + 1
	return rand.Intn(max-min) + min
}

// GeneratePassword 生成加密密码
func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 检查密码是否正确
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
