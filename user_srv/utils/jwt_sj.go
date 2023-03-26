package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"talkon_srvs/user_srv/global"
)

type Token struct {
	Token     string
	SessionId string
	Expire    int64
	UserUid   string
	Platform  string
}

type MyCustomClaims struct {
	Platform  string
	SessionId string
	UserUid   string
	jwt.RegisteredClaims
}

// CreateCustomToken 创建自定义的 Token
// userUid 用户的UUID
// platform 请求的平台
// access 是否将token的SessionId保存到Redis
// expireSe 有效时长（秒）
// subject token的主题
// audience token的观众（token可以使用的平台数组）
func CreateCustomToken(userUid string, platform string,
	access bool, expireSe int64, subject string, audience ...string) (newToken *Token, err error) {
	var jwtConfig = global.ServerConfig.JwtInfo
	myClaims := MyCustomClaims{
		Platform:  platform,
		SessionId: NewUUIDv4MD5Str(),
		UserUid:   userUid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireSe) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jwtConfig.Issuer,
			Subject:   subject,
			ID:        NewUUIDv4Str(),
			Audience:  audience,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenStr, err := token.SignedString([]byte(jwtConfig.Key))
	newToken.Platform = platform
	newToken.Token = tokenStr
	newToken.Expire = myClaims.ExpiresAt.Unix()
	newToken.UserUid = userUid
	newToken.SessionId = myClaims.SessionId
	return newToken, err
}

// CreateConfigToken 根据配置文件创建一个新的 Token
// userUid 用户的id
func CreateConfigToken(userUid string) (newToken *Token, err error) {
	var jwtConfig = global.ServerConfig.JwtInfo
	myClaims := MyCustomClaims{
		Platform:  jwtConfig.Platform,
		SessionId: NewUUIDv4MD5Str(),
		UserUid:   userUid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConfig.Expires) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jwtConfig.Issuer,
			Subject:   jwtConfig.Subject,
			ID:        NewUUIDv4Str(),
			Audience:  []string{jwtConfig.Audience},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenStr, err := token.SignedString([]byte(jwtConfig.Key))
	newToken = new(Token)
	newToken.Platform = myClaims.Platform
	newToken.Token = tokenStr
	newToken.Expire = myClaims.ExpiresAt.Unix()
	newToken.UserUid = myClaims.UserUid
	newToken.SessionId = myClaims.SessionId
	return newToken, err
}

// ParseJwt 解析jwt
func ParseJwt(tokenStr string) (*jwt.Token, error) {
	var key string
	key = global.ServerConfig.JwtInfo.Key
	parseTk, err := jwt.Parse(tokenStr, func(j *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return parseTk, err
}
