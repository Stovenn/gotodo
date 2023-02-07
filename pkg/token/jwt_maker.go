package token

import (
	"fmt"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func (J *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (J *JWTMaker) VerifyToken(token string) (*Payload, error) {
	//TODO implement me
	panic("implement me")
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}
