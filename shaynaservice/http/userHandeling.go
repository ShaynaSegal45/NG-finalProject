package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice/user"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	headerValueContentTypeJSON = "application/json"
	headerKeyContentType       = "Content-Type"
)

type UserService interface {
	SignUp(tx context.Context, user user.User) (string, error)
	SignIn(tx context.Context, user user.User) (user.Profile, error)
	UpdateProfile(tx context.Context, profile user.Profile) (string, error)
	UpdatePassword(tx context.Context, usr user.User) (string, error)
	GetAllUsers(tx context.Context) ([]user.User, error)
	FilterByGender(ctx context.Context, gender string) ([]user.Profile, error)
}

//
//type UserSignUpRequest struct {
//	UserName string
//	Password string
//}

func AddUserRoutes(router *httprouter.Router, s UserService) {
	signUpHandler := makeUserSignUpHandler(s)
	signInHandler := makeUserSignInHandler(s)
	updateHandler := makeUserUpdateProfileHandler(s)
	updatePasswordHandler := makeUserUpdatePasswordHandler(s)
	getAllUsersHandler := makeGetAllUsersHandler(s)
	pingHandler := makepingHandler()
	getUserByGenderHandler := makeGetUserByGender(s)

	router.Handle(http.MethodGet, "/ping", pingHandler)
	router.Handle(http.MethodGet, "/users", getAllUsersHandler)
	router.Handle(http.MethodGet, "/users/:gender", getUserByGenderHandler)
	router.Handle(http.MethodPost, "/sign-up", signUpHandler)
	router.Handle(http.MethodPost, "/sign-in", signInHandler)
	router.Handle(http.MethodPatch, "/my-user/profile", updateHandler)
	router.Handle(http.MethodPatch, "/my-user/password", updatePasswordHandler)
}

// private
func makepingHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		_, err := fmt.Fprint(w, "PONG")
		if err != nil {
			errMessage := fmt.Sprintf("error writing response: %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	}
}
func makeUserSignUpHandler(s UserService) httprouter.Handle {
	fmt.Println("reached http1")
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// decode http request
		fmt.Println("reached http")
		u := &user.User{}
		err := json.NewDecoder(r.Body).Decode(u)

		if err != nil {
			errMessage := fmt.Sprintf("error read from body: %v", err)
			http.Error(w, errMessage, http.StatusBadRequest)
			return
		}
		//
		//userName := params.ByName("userName")
		//password := params.ByName("password")
		//
		//u := user.User{
		//	UserName: userName,
		//	Password: password,
		//}
		_, dbErr := s.SignUp(r.Context(), *u)
		if dbErr != nil {
			panic("make user sign upHandler paniced! " + dbErr.Error())
		}
		formatted := formatUserName(u.UserName)
		encodeJSON(w, formatted)
	}
}
func makeUserUpdateProfileHandler(s UserService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// decode http request
		profile := &user.Profile{}
		err := json.NewDecoder(r.Body).Decode(profile)
		if err != nil {
			errMessage := fmt.Sprintf("error read from body: %v", err)
			http.Error(w, errMessage, http.StatusBadRequest)
			return
		}

		_, dbErr := s.UpdateProfile(r.Context(), *profile)
		if dbErr != nil {
			panic("make user update profile upHandler paniced! " + dbErr.Error())
			return
		}
		return
	}
}
func makeUserUpdatePasswordHandler(s UserService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// decode http request
		u := &user.User{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			errMessage := fmt.Sprintf("error read from body: %v", err)
			http.Error(w, errMessage, http.StatusBadRequest)
			return
		}
		_, dbErr := s.UpdatePassword(r.Context(), *u)
		if dbErr != nil {
			panic("make user update profile upHandler paniced! " + dbErr.Error())
		}

	}
}
func makeUserSignInHandler(s UserService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// decode http request
		//userName := params.ByName("userName")
		//password := params.ByName("password")
		//
		//u := user.User{
		//	UserName: userName,
		//	Password: password,
		u := &user.User{}
		err := json.NewDecoder(r.Body).Decode(u)

		if err != nil {
			errMessage := fmt.Sprintf("error read from body: %v", err)
			http.Error(w, errMessage, http.StatusBadRequest)
			return
		}

		profile, err := s.SignIn(r.Context(), *u)
		if err != nil {
			panic("make sign in Handler paniced!  " + err.Error())
		}
		//formatted:=formatGetUserResponse(username)
		formatted := formatGetProfileResponse(profile)
		encodeJSON(w, formatted)
	}
}

func makeGetAllUsersHandler(s UserService) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Write([]byte("SUCCESS"))
		users, err := s.GetAllUsers(r.Context())
		if err != nil {
			// TODO(oren): encode error (don't panic)
			panic("make get all users Handler paniced!  " + err.Error())
		}

		formatted := formatGetAllUsersResponse(users)
		encodeJSON(w, formatted)

	}
}

func makeGetUserByGender(s UserService) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		g := &user.Gender{}
		err := json.NewDecoder(r.Body).Decode(g)

		if err != nil {
			errMessage := fmt.Sprintf("error read from body: %v", err)
			http.Error(w, errMessage, http.StatusBadRequest)
			return
		}

		w.Write([]byte("SUCCESS"))
		result, err := s.FilterByGender(r.Context(), g.Gender)
		if err != nil {
			// TODO(oren): encode error (don't panic)
			panic("make get all users Handler paniced!  " + err.Error())
		}

		formatted := formatGetAllProfileResponse(result)
		encodeJSON(w, formatted)

	}
}

func encodeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set(headerKeyContentType, headerValueContentTypeJSON)
	//w.Write([]byte("SUCCESS"))
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		// TODO: encode error, don't panic
		panic("aaaaa " + err.Error())
	}
	_, err = w.Write(jsonBytes)
	if err != nil {
		// TODO: log error
	}

	w.WriteHeader(http.StatusTeapot)
}
