package v1

import (
	"fmt"
	"net/http"

	"github.com/gethinyan/sso/models"
	"github.com/gethinyan/sso/pkg/util"
	"github.com/gin-gonic/gin"
)

// GetUserRequest 获取用户信息请求参数
// swagger:parameters GetUserRequest
type GetUserRequest struct {
	// Required: true
	// in: path
	ID uint `json:"id"`
}

// GetUserResponse 获取用户信息响应参数
// swagger:response GetUserResponse
type GetUserResponse struct {
	// in: body
	Body struct {
		// 响应信息
		Message string `json:"message"`
		// 用户信息
		Data models.UserResponseBody `json:"data"`
	}
}

// GetUserInfo 获取用户信息
// GetUserInfo swagger:route GET /users/{id} GetUserRequest
//
// 获取用户信息
//
//     Schemes: http, https
//
//     Responses:
//       200: GetUserResponse
func GetUserInfo(c *gin.Context) {
	user, err := models.GetUserByID(util.UID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "获取用户信息失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.ConvertToResponse()})
}

// UpdateUserRequest 修改用户请求参数
// swagger:parameters UpdateUserRequest
type UpdateUserRequest struct {
	// in: body
	Body models.UserRequestBody
}

// UpdateUserResponse 修改用户响应参数
// swagger:response UpdateUserResponse
type UpdateUserResponse struct {
	// in: body
	Body struct {
		// 响应信息
		Message string `json:"message"`
		// 用户信息
		Data models.UserResponseBody `json:"data"`
	}
}

// UpdateUser 更新用户信息
// UpdateUser swagger:route PUT /users/{id} UpdateUserRequest
//
// 更新用户信息
//
//     Schemes: http, https
//
//     Responses:
//       200: UpdateUserResponse
func UpdateUser(c *gin.Context) {
	var request models.UserRequestBody
	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	user := models.User{
		ID:        util.UID,
		UserName:  request.UserName,
		Phone:     request.Phone,
		Password:  request.Password,
		Email:     request.Email,
		Address:   request.Address,
		Gender:    request.Gender,
		Birth:     request.Birth,
		AvatarURL: request.AvatarURL,
	}
	if err := models.SaveUser(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.ConvertToResponse()})
}
