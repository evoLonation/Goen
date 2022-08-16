package token

import (
	"time"
)

// Maker 用来管理token
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
