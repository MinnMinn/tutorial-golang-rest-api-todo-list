package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"social-todo-list/common"
	"social-todo-list/component/tokenProvider"
	"time"
)

type jwtProvider struct {
	prefix string
	secret string
}

func NewTokenJWTProvider(prefix string, secret string) *jwtProvider {
	return &jwtProvider{prefix: prefix, secret: secret}
}

type myClaims struct {
	Payload common.TokenPayload `json:"payload"`
	jwt.StandardClaims
}
type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *token) GetToken() string {
	return t.Token
}

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

func (j *jwtProvider) Generate(data tokenProvider.TokenPayload, expiry int) (tokenProvider.Token, error) {
	// generate the JWT
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		common.TokenPayload{
			UId:   data.UserId(),
			URole: data.Role(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenProvider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenProvider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenProvider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)

	if !ok {
		return nil, tokenProvider.ErrInvalidToken
	}

	// return the token
	return claims.Payload, nil
}
