package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gethinyan/sso-server/pkg/redis"
	"github.com/gethinyan/sso-server/pkg/util"
)

// Args 请求参数结构
type Args struct {
	JsonWebToken string
}

// Reply 响应参数结构
type Reply struct {
	ID       uint
	Email    string
	NewToken string
}

// Auth ...
type Auth int

// CheckToken 注册函数
func (t *Auth) CheckToken(ctx context.Context, args *Args, reply *Reply) error {
	fmt.Println(args)
	if args.JsonWebToken == "" {
		return errors.New("未传入token")
	}
	// 判断 token 是否在黑名单里面
	tokenExists := redis.RedisClient.HExists("tokenblacklist", args.JsonWebToken).Val()
	if tokenExists == true {
		return errors.New("token在黑名单里")
	}
	user, err := util.ParseToken(args.JsonWebToken)
	fmt.Println(user)
	if err != nil {
		return err
	}
	reply.ID = user.ID
	reply.Email = user.Email
	// 判断 token 是否有效
	if time.Now().Unix() > user.ValidBefore {
		newToken, err := util.GenerateToken(user.ID, user.Email)
		fmt.Println(newToken, err)
		if err != nil {
			return err
		}
		reply.NewToken = newToken
		// 把老 token 加入黑名单
		redis.RedisClient.HSet("tokenblacklist", args.JsonWebToken, true)
	}

	return nil
}
