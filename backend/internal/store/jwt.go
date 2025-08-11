package store

import (
	"time"

	"marketplace/backend/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	conf config.JWTConfig
}

func NewJWT(conf config.JWTConfig) *JWT { return &JWT{conf: conf} }

func (j *JWT) SignAccess(userID int64, role string) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(j.conf.AccessTTL)
	claims := jwt.MapClaims{
		"sub": userID,
		"role": role,
		"iat": now.Unix(),
		"exp": exp.Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(j.conf.Secret))
	return s, exp, err
}
