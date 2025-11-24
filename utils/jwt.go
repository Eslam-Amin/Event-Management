package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}


func ValidateToken(token string) error {
	parsedToken, err :=jwt.Parse(token, func(token *jwt.Token)(interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.ErrTokenInvalidClaims
	}

	return nil


	// email := claims["email"].(string) 
	// userId := claims["userId"].(int64)
	
	// return nil
}