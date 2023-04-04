package authManager

import "time"

type UserId uint64

type TokenPayload struct {
	Id  UserId    `valid:"userId, required"`
	Exp time.Time `valid:"expired, required"`
}

const (
	idTitle      = "id"
	expiresTitle = "expired"
)

func NewTokenPayload(id UserId) *TokenPayload {
	return &TokenPayload{
		Id: id,
	}
}

type AuthManager interface {
	CreateToken(payload *TokenPayload) (string, error)
	ParseToken(token string) (*TokenPayload, error)
	GetEpiryTime() time.Duration
}

func TokenPayloadToMap(payload TokenPayload) map[string]interface{} {
	return map[string]interface{}{
		idTitle:      payload.Id,
		expiresTitle: payload.Exp,
	}
}

func MapToTokenPayload(payloadMap map[string]interface{}) *TokenPayload {
	expStr, ok := payloadMap[expiresTitle].(string)
	if !ok {
		return nil
	}
	exp, _ := time.Parse(time.RFC3339, expStr)
	return &TokenPayload{
		Id:  UserId(payloadMap[idTitle].(float64)),
		Exp: exp,
	}
}
