package token

import (
	"fmt"
	"strings"
	"time"
)

type Token struct {
	AccessToken     string
	AccessTokenType string
}

func FromCookieValue(value string) (*Token, error) {
	parts := strings.SplitN(value, " ", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid cookie value")
	}
	return &Token{
		AccessToken:     parts[1],
		AccessTokenType: parts[0],
	}, nil
}

func (token *Token) ToCookieValue() string {
	return token.AccessTokenType + " " + token.AccessToken
}

type Payload struct {
	Iss string
	Sub string
	Iat int64
	Exp int64
}

func NewPayload(sub string) *Payload {
	return &Payload{
		Iss: "zvax-auth",
		Sub: sub,
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(time.Hour * 24).Unix(),
	}
}

func (payload *Payload) ToClaimsMap() map[string]interface{} {
	return map[string]interface{}{
		"iss": payload.Iss,
		"sub": payload.Sub,
		"iat": payload.Iat,
		"exp": payload.Exp,
	}
}

func (payload *Payload) FromClaimsMap(claims map[string]interface{}) {
	iat := claims["iat"].(float64)
	exp := claims["exp"].(float64)
	iatInt := int64(iat)
	expInt := int64(exp)

	payload.Iss = claims["iss"].(string)
	payload.Sub = claims["sub"].(string)
	payload.Iat = iatInt
	payload.Exp = expInt
}
