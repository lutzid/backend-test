package auth

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds specific field as a claim to the token
type JwtClaim struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	jwt.RegisteredClaims
}

type AdditionalJwtClaim struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

// GenerateToken generates a jwt token
func (j *JwtWrapper) GenerateToken(claim AdditionalJwtClaim) (signedToken string, err error) {
	expiredTime := time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaim{
		Name:      claim.Name,
		Phone:     claim.Phone,
		Role:      claim.Role,
		CreatedAt: claim.CreatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	})

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

// ValidateToken validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt.Before(time.Now().Local()) {
		err = errors.New("JWT is expired")
		return
	}

	return

}
