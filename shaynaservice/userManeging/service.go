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
	AlterProfile(ctx context.Context, profile user.Profile) error
	AlterPassword(ctx context.Context, user user.User) error
	GetProfilesByGender(ctx context.Context, gender string) ([]user.Profile, error)
	UserProfile(ctx context.Context, userName string) (user.Profile, error)

}

type service struct {
	userRepo UserRepo
}

func NewService(user UserRepo) *service {
	return &service{
		userRepo: user,
	}
}

func (s *service) SignUp(tx context.Context, usr user.User) (string, error) {
	//if checkPassword(usr.Password)

	err := s.userRepo.InsertUser(tx, usr)
	if err != nil {
		return "", err
	}

	return usr.UserName, nil

}
func (s *service) UpdateProfile(tx context.Context, p user.Profile) (string, error) {
	err := s.userRepo.AlterProfile(tx, p)
	if err != nil {
		return "", fmt.Errorf("Update user: %w", err)
	}

	return p.UserName, nil
}
func (s *service) UpdatePassword(tx context.Context, usr user.User) (string, error) {
	err := s.userRepo.AlterPassword(tx, usr)
	if err != nil {
		return "", fmt.Errorf("Update user: %w", err)
	}

	return usr.UserName, nil
}
func (s *service) SignIn(ctx context.Context, usr user.User) (user.Profile, error) {

	results, err := s.userRepo.LoginUser(ctx, usr.UserName)
	if err != nil {
		return user.Profile{}, err
	}

	if results.Password != usr.Password {
		return user.Profile{}, err
	}
//return profile data
	profile, err := s.userRepo.UserProfile(ctx, usr.UserName)
	if err != nil {
		return user.Profile{}, err
	}
	return profile, nil

}

func (s *service) GetAllUsers(ctx context.Context) ([]user.User, error) {
	users, err := s.userRepo.FindAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateDuck: %w", err)
	}

	return users, nil
}

func (s *service) FilterByGender(ctx context.Context, gender string) ([]user.Profile, error) {
	users, err := s.userRepo.GetProfilesByGender(ctx, gender)
	if err != nil {
		return nil, fmt.Errorf("CreateDuck: %w", err)
	}

	return users, nil
}
