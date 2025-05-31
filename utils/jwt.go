package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// var jwtKey = []byte(os.Getenv("JWT_SECRET"))
var jwtKey = []byte("bb557692062675bf9dc8f4711ae2a57f55dcb28d1443d0f3ee377c36be15fbdc43a7f8f5f201f9772b0022304de9fd81ad3ad011b47fa2f905ccab1fd7ff1a18")
func GenerateJWT(email string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["email"].(string), nil
	}

	return "", err
}
