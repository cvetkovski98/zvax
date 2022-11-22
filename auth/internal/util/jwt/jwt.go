package jwtutil

import (
	"errors"

	"github.com/cvetkovski98/zvax-auth/internal/token"
	"github.com/golang-jwt/jwt/v4"
)

func Generate(payload *token.Payload, secret string) (*token.Token, error) {
	claims := jwt.MapClaims(payload.ToClaimsMap())
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &token.Token{
		AccessToken:     tokenString,
		AccessTokenType: "Bearer",
	}, nil
}

func ParseVerify(t *token.Token, secret string) (*token.Payload, error) {
	jwtToken, err := jwt.Parse(t.AccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, err
	}
	payload := new(token.Payload)
	payload.FromClaimsMap(claims)
	return payload, nil
}
