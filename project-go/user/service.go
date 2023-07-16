package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	var user User
	user.Name = input.Name
	user.Email = input.Email
	user.NoHandphone = input.NoHandphone

	checkUserEmail, err := s.repository.FindUserByEmail(user.Email)
	if err != nil {
		return checkUserEmail, err
	}

	if checkUserEmail.ID != 0 {
		return checkUserEmail, errors.New("email ini sudah terdaftar")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.CreateUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("data user tidak ditemukan berdasarkan email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return false, err
	}

	// Jika email tidak ditemukan dalam database maka email available
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindUserByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("data user tidak ditemukan berdasarkan ID")
	}

	return user, nil
}
