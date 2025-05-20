package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)


type Service interface {
	GeneratedToken(userID uuid.UUID)(string, error)
	ValidateToken(encodedToken string)(*jwt.Token, error)
}

type jwtService struct {
    
}

func NewService() *jwtService {
	return &jwtService{}
}
var SECRET_KEY = []byte(os.Getenv("KEY_SECRET"))

func (s *jwtService) GeneratedToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"id_user": userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}


func (s *jwtService) ValidateToken(encodedToken string)(*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {		
		_, ok := t.Method.(*jwt.SigningMethodHMAC) 

		if !ok {
			fmt.Println(!ok, "error not ok")
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		fmt.Println(err.Error(), "error token parse")
		return token, err
	}

	fmt.Println(token, "hasil token parse")
	return token, nil
}