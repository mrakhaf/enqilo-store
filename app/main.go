package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authStaffHandler "github.com/mrakhaf/enqilo-store/domain/auth-staff/delivery/http"
	authRepository "github.com/mrakhaf/enqilo-store/domain/auth-staff/repository"
	authUsecase "github.com/mrakhaf/enqilo-store/domain/auth-staff/usecase"
	customerHandler "github.com/mrakhaf/enqilo-store/domain/customer/delivery/http"
	customerRepository "github.com/mrakhaf/enqilo-store/domain/customer/repository"
	customerUsecase "github.com/mrakhaf/enqilo-store/domain/customer/usecase"
	productHandler "github.com/mrakhaf/enqilo-store/domain/product/delivery/http"
	productRepository "github.com/mrakhaf/enqilo-store/domain/product/repository"
	productUsecase "github.com/mrakhaf/enqilo-store/domain/product/usecase"
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
	group := e.Group("/v1")

	productGroup := e.Group("/v1")
	productGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
	}))

	//
	formatResponse := formatJson.NewResponse()
	jwtAccess := jwt.NewJWT()

	//auth-staff
	authRepo := authRepository.NewRepository(database)
	authUsecase := authUsecase.NewUsecase(authRepo, jwtAccess)
	authStaffHandler.AuthHandler(group, authUsecase, authRepo, formatResponse)

	//customer
	customerRepo := customerRepository.NewRepository(database)
	customerUsecase := customerUsecase.NewUsecase(customerRepo)
	customerHandler.CustomerHandler(group, customerUsecase, customerRepo, formatResponse)

	//product
	productRepo := productRepository.NewRepository(database)
	productUsecase := productUsecase.NewUsecase(customerRepo, productRepo)
	productHandler.ProductHandler(productGroup, productUsecase, productRepo, formatResponse)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
