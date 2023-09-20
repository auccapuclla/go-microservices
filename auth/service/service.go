package service

import (
	"context"
	"microservices/auth/models"
	"microservices/auth/repository"
	"microservices/auth/validators"
	"microservices/pb"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type authService struct {
	usersRepository repository.UsersRepository
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(usersRepository repository.UsersRepository) pb.AuthServiceServer {
	return &authService{usersRepository: usersRepository}
}

func (s *authService) SignUp(ctx context.Context, req *pb.User) (*pb.User, error) {
	err := validators.ValidateSignUp(req)
	if err != nil {
		return nil, err
	}
	foundUser, err := s.usersRepository.GetByEmail(req.Email)
	if err == mongo.ErrNoDocuments {
		user := new(models.User)
		user.FromProtoBuffer(req)
		err := s.usersRepository.Save(user)
		if err != nil {
			return nil, err
		}
		return user.ToProtoBuffer(), nil
	}
	if foundUser == nil {
		return nil, err
	}
	return nil, validators.ErrEmptyEmail
}

func (s *authService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	userId, _ := primitive.ObjectIDFromHex(req.Id)
	foundUser, err := s.usersRepository.GetById(userId)
	if err != nil {
		return nil, err
	}
	return foundUser.ToProtoBuffer(), nil

}

func (s *authService) ListUsers(*pb.ListUsersRequest, pb.AuthService_ListUsersServer) error {
	users, err := s.usersRepository.GetAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		// err := stream.Send(user.ToProtoBuffer())
		// if err != nil {
		// 	return nil, err
		// }
		println(user.Name)
	}
	return nil
}

func (s *authService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if !primitive.IsValidObjectID(req.Id) {
		return nil, validators.ErrInvalidId
	}
	id, _ := primitive.ObjectIDFromHex(req.Id)
	userFound, err := s.usersRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, validators.ErrEmptyName
	}
	if req.Name == userFound.Name {
		return userFound.ToProtoBuffer(), nil
	}

	userFound.Name = req.Name
	userFound.Updated = time.Now()
	err = s.usersRepository.Update(userFound)
	return userFound.ToProtoBuffer(), err
}

func (s *authService) DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteUserResponse, error) {
	if !primitive.IsValidObjectID(req.Id) {
		return nil, validators.ErrInvalidId
	}
	err := s.usersRepository.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Id: req.Id}, nil
}
