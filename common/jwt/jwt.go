package jwt

import (
	"context"
	"go-mq/common/redis"
	"go-mq/common/util"
	"go-mq/common/xerr"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type Jwt struct {
	logx.Logger
	Key string
	Ctx context.Context
}

const (
	EmailLoginTokenType = "email_login"
	ResetPwdTokenType   = "reset_password"
	LoginTokenType      = "login"
)

func NewJwt(ctx context.Context, key string) *Jwt {
	return &Jwt{
		Key:    key,
		Ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (p *Jwt) GetJwtToken(user, username string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_id"] = user
	claims["username"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(p.Key))
}

// ParseToken 解析token
func (p *Jwt) ParseToken(tokenString string) (string, error) {
	InvalidError := util.ReturnError(xerr.InvalidToken)
	if p.CheckBlack(tokenString) {
		p.Logger.Errorf("Jwt CheckBlack: 无效的token token:%s", tokenString)
		return "", InvalidError
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			p.Logger.Errorf("Jwt  unexpected signing method: %v", token.Header["alg"])
			return nil, InvalidError
		}
		// Return the secret key for verification
		return []byte(p.Key), nil
	})
	if err != nil {
		p.Logger.Errorf("Jwt Parse error : %+v", err)
		return "", InvalidError
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["user"].(string)
		// 将过期时间转化为时间戳
		exp := time.Unix(int64(claims["exp"].(float64)+claims["iat"].(float64)), 0)

		// 判断是否过期
		if !time.Now().Before(exp) {
			p.Logger.Errorf("Jwt Parse error : Token has expired")

			return "", InvalidError
		}

		return user, nil
	}
	p.Logger.Errorf("Jwt token.Claims error : Handle invalid token")
	return "", InvalidError

}

// AddBlack 加入到黑名单
func (p *Jwt) AddBlack(token string) (err error) {
	p.Logger.Infof("Jwt AddBlack token:%s", token)
	redis.Rdb.SAdd(p.Ctx, "logout", token)
	return
}

// CheckBlack 检查token是否在黑名单
func (p *Jwt) CheckBlack(token string) bool {
	isExist, err := redis.Rdb.SIsMember(p.Ctx, "logout", token).Result()
	if err == nil && isExist {
		return true
	}
	return false
}

func (p *Jwt) GetEmailToken(user, username string, iat, seconds int64, tokenType string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_id"] = user
	claims["type"] = tokenType
	claims["username"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(p.Key))
}

// ParseOriginalToken 解析token
func (p *Jwt) ParseOriginalToken(tokenString string) (jwt.MapClaims, error) {
	InvalidError := util.ReturnError(xerr.InvalidToken)
	if p.CheckBlack(tokenString) {
		p.Logger.Errorf("Jwt CheckBlack: 无效的token token:%s", tokenString)
		return nil, InvalidError
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			p.Logger.Errorf("Jwt  unexpected signing method: %v", token.Header["alg"])
			return nil, InvalidError
		}
		// Return the secret key for verification
		return []byte(p.Key), nil
	})
	if err != nil {
		p.Logger.Errorf("Jwt Parse error : %+v", err)
		return nil, InvalidError
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	p.Logger.Errorf("Jwt token.Claims error : Handle invalid token")
	return nil, InvalidError

}
