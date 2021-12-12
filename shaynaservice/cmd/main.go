package main

import (
	userhttp "github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice/http"
	"github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice/sql"
	"github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice/userManeging"

	"fmt"

	"net/http"
	"os"

	"github.com/RevitalS/someone-to-run-with-app/backend/foundation/config"
	"github.com/RevitalS/someone-to-run-with-app/backend/foundation/nextlog"
	"github.com/RevitalS/someone-to-run-with-app/backend/foundation/nextsql"
)

const (
	port = ":9011"
)

func main() {
	fmt.Println("Hello World")

	configurator := config.NewConfigurator("config/development")

	// loggerator (loggers factory)
	loggerator := nextlog.NewLoggerator("debug", os.Stdout)
	mainLogger := loggerator.NewLogger("main")

	// sql
	sqlConfig := configurator.SQLConfig()
	fmt.Println("sql confit")
	sqlLogger := loggerator.SQLLogger()
	sqlDB, err := nextsql.NewDB(sqlConfig, sqlLogger)
	if err != nil {
		mainLogger.Error(err, "failed to open db connection")
	}

	// user repo
	userRepo := sql.NewUserRepo(sqlDB)

	// user managing
	userService := userManeging.NewService(userRepo)

	fmt.Println("b4 http router")
	// starting the http server
	router := userhttp.NewRouter()

	userhttp.AddUserRoutes(router, userService)

	fmt.Println("ater user http")
	mainLogger.Info(fmt.Sprintf("start server on port: %s. if development - http://localhost%s", port, port))
	if err := http.ListenAndServe(port, router); err != nil {
		mainLogger.Error(err, "failed http listening and serving")
	}

}
