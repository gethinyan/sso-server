package models

import (
	"fmt"
	"time"

	"github.com/gethinyan/sso-server/pkg/util"
	"github.com/jinzhu/gorm"
)

// TableName 指定用户表表名
func (User) TableName() string {
	return "users"
}

// User 定义用户表对应的结构
type User struct {
	ID        uint      `json:"id" gorm:"primary_key;not null;auto_increment"`
	Phone     string    `json:"phone" gorm:"type:char(11);not null;default:''"`
	Email     string    `json:"email" gorm:"size:50;not null;default:''"`
	UserName  string    `json:"user_name" gorm:"size:50;not null;default:''"`
	Password  string    `json:"password" gorm:"type:char(60);not null;default:''"`
	Address   string    `json:"address" gorm:"size:200;not null;default:''"`
	Gender    int8      `json:"gender" gorm:"not null;default:1"`
	Birth     time.Time `json:"birth" gorm:"not null;type:date;default:'1970-01-01'"`
	AvatarURL string    `json:"avatar_url" gorm:"size:200;not null;default:''"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:current_timestamp"`
}

// UserRequestBody 用户请求参数
// swagger:parameters UserRequestBody
type UserRequestBody struct {
	// 手机号
	// Required: true
	Phone string `json:"phone" binding:"required"`
	// 邮箱地址
	// Required: true
	Email string `json:"email" binding:"required,email"`
	// 用户名
	// Required: true
	UserName string `json:"user_name" binding:"required"`
	// 密码
	// Required: true
	Password string `json:"password" binding:"required"`
	// 地址
	Address string `json:"address"`
	// 性别（0：女；1：男）
	// Required: true
	Gender int8 `json:"gender" binding:"required"`
	//生日
	Birth time.Time `json:"birth"`
	// 头像地址
	AvatarURL string `json:"avatar_url"`
	// 验证码（注册时必填）
	Code int `json:"code"`
}

// UserResponseBody 用户响应参数
// swagger:parameters UserResponseBody
type UserResponseBody struct {
	// 用户 ID
	ID uint `json:"id"`
	// 手机号
	Phone string `json:"phone"`
	// 邮箱地址
	Email string `json:"email"`
	// 用户名
	UserName string `json:"user_name"`
	// 地址
	Address string `json:"address"`
	// 性别（0：女；1：男）
	Gender int8 `json:"gender"`
	// 生日
	Birth time.Time `json:"birth"`
	// 头像地址
	AvatarURL string `json:"avatar_url"`
}

// GetUserByID 通过 ID 获取用户信息
func GetUserByID(id uint) (user User, err error) {
	if err = dbConn.Where(User{ID: id}).Find(&user).Error; err != nil {
		return
	}
	return
}

// GetUserByEmail 通过邮箱获取用户信息
func GetUserByEmail(email string) (user User, err error) {
	if err = dbConn.Where(User{Email: email}).Find(&user).Error; err != nil {
		return
	}
	return
}

// BeforeSave ...
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if user.Password, err = util.GeneratePassword(user.Password); err != nil {
		return
	}
	scope.SetColumn("Password", user.Password)
	return
}

// SaveUser 保存用户信息，包括创建/更新
func SaveUser(user *User) error {
	if err := dbConn.Save(user).Error; err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}

// ConvertToResponse 转换为响应参数
func (user *User) ConvertToResponse() UserResponseBody {
	response := UserResponseBody{
		ID:        user.ID,
		Phone:     user.Phone,
		Email:     user.Email,
		UserName:  user.UserName,
		Address:   user.Address,
		Gender:    user.Gender,
		Birth:     user.Birth,
		AvatarURL: user.AvatarURL,
	}
	return response
}
