package jwtManager

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/tools/authManager"
	"time"
)

type JwtManager struct {
	key         []byte
	method      jwt.SigningMethod
	expDuration time.Duration
}

func NewJwtManager(jwtS string) *JwtManager {
	methodObj := jwt.GetSigningMethod("HS256")
	t, _ := time.ParseDuration("168h")
	if methodObj == nil {
		return nil
	}
	return &JwtManager{
		key:         []byte(jwtS),
		method:      methodObj,
		expDuration: t,
	}
}

func (manager *JwtManager) GetEpiryTime() time.Duration {
	return manager.expDuration
}

func (manager *JwtManager) CreateToken(payload *authManager.TokenPayload) (string, error) {
	payload.Exp = time.Now().Add(manager.expDuration)
	token := jwt.NewWithClaims(manager.method, jwt.MapClaims(authManager.TokenPayloadToMap(*payload)))
	if token == nil {
		return "", errors.Errorf("произошла ошибка")
	}
	tokenSigned, err := token.SignedString(manager.key)
	if err != nil {
		return "", errors.Errorf("произошла ошибка")
	}
	return tokenSigned, nil
}

func (manager *JwtManager) ParseToken(token string) (*authManager.TokenPayload, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return manager.key, nil
	})

	if jwtToken == nil || err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, errors.Errorf("произошла ошибка")
	}
	return authManager.MapToTokenPayload(claims), nil
}
