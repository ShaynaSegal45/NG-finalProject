package http

import (
	"github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice/user"
)

func formatGetUserResponse(usr user.User) map[string]interface{} {
	return map[string]interface{}{
		"user": formatUser(usr),
	}
}
func formatGetProfileResponse(profile user.Profile) map[string]interface{} {
	return map[string]interface{}{
		"profile": formatProfile(profile),
	}
}


func formatUserName(userName string) map[string]interface{} {
	return map[string]interface{}{
		"userName": userName,
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
func formatGetAllProfileResponse(profiles []user.Profile) map[string]interface{} {
	formattedProfile := make([]map[string]interface{}, len(profiles))
	for i, p := range profiles {
		formattedProfile[i] = formatProfile(p)
	}

	return map[string]interface{}{
		"profile": formattedProfile,
	}
}

func formatUser(usr user.User) map[string]interface{} {
	return map[string]interface{}{
		"userName": usr.UserName,
		"password": usr.Password,
	}
}
func formatProfile(profile user.Profile) map[string]interface{} {
	return map[string]interface{}{
		"userName": profile.UserName,
		"gender": profile.Gender,
		"age": profile.Age,
	}
}