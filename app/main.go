package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authStaffHandler "github.com/mrakhaf/enqilo-store/domain/auth-staff/delivery/http"
	"github.com/mrakhaf/enqilo-store/domain/auth-staff/repository"
	"github.com/mrakhaf/enqilo-store/domain/auth-staff/usecase"
	"github.com/mrakhaf/enqilo-store/shared/common"
	formatJson "github.com/mrakhaf/enqilo-store/shared/common/json"
	"github.com/mrakhaf/enqilo-store/shared/common/jwt"
	"github.com/mrakhaf/enqilo-store/shared/config/database"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = common.NewValidator()

	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatal(err)
	}

	//db config
	database, err := database.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Enqilo Store!")
	})

	//create group
	userGroup := e.Group("/v1/staff")

	// catGroup := e.Group("/v1/cat")
	// catGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningMethod: "HS256",
	// 	SigningKey:    []byte("secret"),
	// }))

	//
	formatResponse := formatJson.NewResponse()
	jwtAccess := jwt.NewJWT()

	//auth-staff
	authRepo := repository.NewRepository(database)
	authUsecase := usecase.NewUsecase(authRepo, jwtAccess)
	authStaffHandler.AuthHandler(userGroup, authUsecase, authRepo, formatResponse)

	//cat
	// catHandler.CatHandler(catGroup, formatResponse, jwtAccess)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
