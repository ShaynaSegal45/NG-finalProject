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

type UserRepo struct {
	db nextsql.DB
}

func NewUserRepo(db nextsql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) InsertUser(ctx context.Context, user user.User) error {
	q := `
		INSERT into users (userName, password)
		VALUES (?, ?)`

	_, err := r.db.Exec(context.Background(), q, user.UserName, user.Password)
	if err != nil {
		return fmt.Errorf("failed to insert user to db: %w", err)
	}

	return nil
}


func (r *UserRepo) LoginUser(ctx context.Context, userName string ) (user.User, error) {
	q := `
		SELECT *
		FROM users
		WHERE userName = ?
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
