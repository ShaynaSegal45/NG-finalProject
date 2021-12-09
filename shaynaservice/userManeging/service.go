package userManeging

import (
	"context"
	"fmt"
	"github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice/user"
)

type UserRepo interface {
	InsertUser(ctx context.Context, user user.User) error
	LoginUser(ctx context.Context, userName string) (user.User, error)
	FindAllUsers(ctx context.Context) ([]user.User, error)
	//PingUser(ctx context.Context, user user.User) error

}

type service struct {
	userRepo UserRepo
}

func NewService(user UserRepo) *service {
	return &service{
		userRepo: user,
	}
}

func (s *service) SignUp(tx context.Context, usr user.User) (str string, e error) {
	//if checkPassword(usr.Password)
	err := s.userRepo.InsertUser(tx, usr)
	if err != nil {
		return "", err
	}
   str= usr.UserName
	return str, nil

}
func (s *service) SignIn(ctx context.Context, usr user.User) (string, error) {

	results, err := s.userRepo.LoginUser(ctx, usr.UserName)
	if err != nil {
		return "", err
	}
	if results.Password != usr.Password {
		return "", err
	}
	return usr.UserName, nil

}

func (s *service) GetAllUsers(ctx context.Context) ([]user.User, error) {
	users, err := s.userRepo.FindAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateDuck: %w", err)
	}

	return users, nil
}

//
//func (s *service) SignPing(tx context.Context) ( string, error) {
//	_, err := fmt.Fprint(w, "PONG")
//	if err != nil {
//		errMessage := fmt.Sprintf("error writing response: %v", err)
//		http.Error(w, errMessage, http.StatusInternalServerError)
//	}
//
//
//	return "nil",nil
//
//}
