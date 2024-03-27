package utils

import "github.com/dgrijalva/jwt-go"

// define custom claims

type CustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// Generate JWT token
func GenerateToken(username, password string) (string, error) {
	claims := CustomClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("gkuyGUYTGUYf76F87^f9FYf67YU))"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Parse and validate JWT token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("gkuyGUYTGUYf76F87^f9FYf67YU"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
