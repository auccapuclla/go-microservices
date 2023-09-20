package validators

import (
	"errors"
	"microservices/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidId     = errors.New("invalid id")
	ErrEmptyName     = errors.New("name is empty")
	ErrEmptyEmail    = errors.New("email is empty")
	ErrEmptyPassword = errors.New("password is empty")
)

func ValidateSignUp(user *pb.User) error {
	if !primitive.IsValidObjectID(user.Id) {
		return ErrInvalidId
	}
	if user.Name == "" {
		return ErrEmptyName
	}
	if user.Email == "" {
		return ErrEmptyEmail
	}
	if user.Password == "" {
		return ErrEmptyPassword
	}
	return nil
}
