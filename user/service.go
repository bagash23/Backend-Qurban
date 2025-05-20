package user

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterUserInput)(User, error)
	LoginUser(input LoginUserInput)(User, error)
	IsEmailAvailable(input CheckEmailInput)(bool, error)
	GetUserByID(id string)(User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}


func (s service) RegisterUser(input RegisterUserInput)(User, error) {
	user := User{}
	user.IDUser = uuid.New()
	user.Username = input.Username
	user.Email = input.Email
	passwordHas, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Role = input.Role
	user.Password = string(passwordHas)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s service) LoginUser(input LoginUserInput)(User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.IDUser == uuid.Nil {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput)(bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.IDUser == uuid.Nil {
		return true, nil
	}

	return false, nil
}

func (s service) GetUserByID(id string)(User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	if user.IDUser == uuid.Nil {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}