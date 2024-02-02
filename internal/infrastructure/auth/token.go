package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type (
	Claims struct {
		CustomerXID string `json:"customer_xid"`
		jwt.StandardClaims
	}

	module struct {
		config Config
	}

	Config struct {
		Secret    string `env:"SECRET"`
		ExpInHour int    `env:"EXP_IN_HOUR"`
	}

	Auth interface {
		GenerateToken(claims Claims) (string, error)
		VerifyToken(tokenString string) (*Claims, error)
	}
)

func New(cfg Config) Auth {
	return &module{config: cfg}

}

func (t *module) GenerateToken(claims Claims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(t.config.ExpInHour)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(t.config.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *module) VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
