package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const minSecretKeyLength = 32

var algorithm = jwt.SigningMethodHS256
var ErrInvalidToken = errors.New("token is invalid")

type JWTMaker struct {
	secretKey string
}

func (m *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(algorithm, payload)
	return jwtToken.SignedString([]byte(m.secretKey))
}

func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {

	// 先检查token的header部分的alg是否符合服务器的alg
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		alg, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || alg != algorithm {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	}
	// 将token字符串转化为token类型的值。
	// keyFunc: 返回verify时用到的key，同时检查token的有效性
	// claims: 用于定义载荷的类型，期间也会用valid函数判断有效性
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeyLength {
		return nil, fmt.Errorf("secret key's length must longer than %d ", minSecretKeyLength)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}
