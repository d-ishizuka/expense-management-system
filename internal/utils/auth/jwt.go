package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrTokenInvalid = errors.New("invalid token")
	secretKey      = []byte("your_secret_key") // 秘密鍵は環境変数などで管理することを推奨
)

// GenerateToken は新しいJWTトークンを生成します。
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 24時間後に期限切れ
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken はJWTトークンを検証します。
func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", ErrTokenInvalid
	}

	userID := claims["sub"].(string)
	return userID, nil
}