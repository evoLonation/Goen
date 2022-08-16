package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var ErrExpiredToken = errors.New("token has expired")

// Payload 有效载荷，拥有方法Valid用来判断是否仍然有效
type Payload struct {
	// uuid 是通用唯一标识符
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		Id:        id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
