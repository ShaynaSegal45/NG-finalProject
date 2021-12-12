package sql

import (
	"context"
	"fmt"
	"github.com/RevitalS/someone-to-run-with-app/backend/foundation/nextsql"

	"github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice/user"
	//"strconv"
)

type userRes struct {
	UserName string `sql:"userName"`
	Password string `sql:"password"`
}

type profileRes struct {
	UserName string `sql:"userName"`
	Gender string `sql:"gender"`
	Age int `sql:"age"`
}


type UserRepo struct {
	db nextsql.DB
}

func NewUserRepo(db nextsql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) InsertUser(ctx context.Context, user user.User) error {

	q1 := `select * from users where userName=?`
	exists, err := r.db.IsExists(ctx, q1, userRes{}, user.UserName)
	if exists {
		return fmt.Errorf("try a new user: %w", err)
	}
	if err != nil{
		return fmt.Errorf("db error: %w", err)
	}
	q := `
		INSERT into users (userName, password)
		VALUES (?, ?)
`

	_, err = r.db.Exec(context.Background(), q, user.UserName, user.Password)
	if err != nil {
		return fmt.Errorf("failed to insert user to db: %w", err)
	}

	return nil
}
func (r *UserRepo) AlterProfile(ctx context.Context, profile user.Profile) error {
	q := `
			UPDATE profiles
			SET gender = ?, age= ?
			WHERE userName = ?
`

	_, err := r.db.Exec(context.Background(), q, profile.Gender, profile.Age, profile.UserName)
	if err != nil {
		return fmt.Errorf("failed to insert user to db: %w", err)
	}

	return nil
}
func (r *UserRepo) AlterPassword(ctx context.Context, user user.User) error {
	q := `
		UPDATE users
		SET password = ?
		WHERE userName = ?;
`

	_, err := r.db.Exec(context.Background(), q, user.Password, user.UserName)
	if err != nil {
		return fmt.Errorf("failed to insert user to db: %w", err)
	}

	return nil
}

func (r *UserRepo) LoginUser(ctx context.Context, userName string) (user.User, error) {
	q1 := `select * from users where userName=?`
	exists, err := r.db.IsExists(ctx, q1, userRes{}, userName)
	if !exists {
		return user.User{},fmt.Errorf("user doesnt exist: %w", err)
	}
	if err != nil{
		return user.User{},fmt.Errorf("db error: %w", err)
	}
	q := `
		SELECT *
		FROM users
		WHERE userName = ? 
		and passwored=?
`

	result, err := r.db.StructQuery(context.Background(), userRes{}, q, userName)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to login user to db: %w", err)
	}
	if len(result) == 0 {
		return user.User{}, fmt.Errorf("didn't find user with this name: ")
		//return user.User{}, fmt.Errorf("didn't find user with name: %s", userName)
	}
	res := result[0]
	userResult := res.(userRes)

	return user.User{
		UserName: userResult.UserName,
		Password: userResult.Password,
	}, nil

}

//
func (r *UserRepo) UserProfile(ctx context.Context, userName string) (user.Profile, error) {
	q := `select * from profiles where userName=?`
	result, err := r.db.StructQuery(context.Background(), profileRes{}, q, userName)
	if err != nil {
		return user.Profile{}, fmt.Errorf("failed to login user to db: %w", err)
	}
	if len(result) == 0 {
		return user.Profile{}, fmt.Errorf("didn't find user with this name: ")
		//return user.User{}, fmt.Errorf("didn't find user with name: %s", userName)
	}
	res := result[0]
	profileResult := res.(profileRes)

	return user.Profile{
		UserName: profileResult.UserName,
		Gender: profileResult.Gender,
		Age: profileResult.Age,
	}, nil

}
//
func (r *UserRepo) FindAllUsers(ctx context.Context) ([]user.User, error) {
	q := `
		SELECT *
		FROM users
`

	results, err := r.db.StructQuery(context.Background(), userRes{}, q)

	if err != nil {
		return nil, fmt.Errorf("failed to query from sql: %w", err)
	}

	users := make([]user.User, len(results))
	for i, res := range results {
		userRes := res.(userRes)
		users[i] = user.User{
			UserName: userRes.UserName,
			Password: userRes.Password,
		}
	}
	return users, nil
}

func (r *UserRepo) GetProfilesByGender(ctx context.Context, gender string) ([]user.Profile, error) {
	q := `
		select * 
		from profiles
		where gender=?
`

	results, err := r.db.StructQuery(context.Background(), userRes{}, q, gender)

	if err != nil {
		return nil, fmt.Errorf("failed to query from sql: %w", err)
	}

	profiles := make([]user.Profile, len(results))
	for i, res := range results {
		profileRes := res.(profileRes)
		profiles[i] = user.Profile{
			UserName: profileRes.UserName,
			Gender: profileRes.Gender,
			Age: profileRes.Age,
		}
	}

	return profiles, nil
}