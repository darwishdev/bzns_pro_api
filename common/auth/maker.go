package auth

import "time"

type UserPayload struct {
	Username string
	UserId   int32
	Duration time.Duration
}
type Maker interface {
	CreateToken(userPayload UserPayload) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
