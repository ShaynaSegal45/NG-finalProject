package http

import (
	"github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice/user"
)

func formatGetUserResponse(usr user.User) map[string]interface{} {
	return map[string]interface{}{
		"user": formatUser(usr),
	}
}

func formatGetAllUsersResponse(users []user.User) map[string]interface{} {
	formattedUsers := make([]map[string]interface{}, len(users))
	for i, usr := range users {
		formattedUsers[i] = formatUser(usr)
	}

	return map[string]interface{}{
		"users": formattedUsers,
	}
}

func formatUser(usr user.User) map[string]interface{} {
	return map[string]interface{}{
		"userName": usr.UserName,
		"password": usr.Password,
	}
}
