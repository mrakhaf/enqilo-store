package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/enqilo-store/domain/auth-staff/interfaces"
	"github.com/mrakhaf/enqilo-store/models/dto"
	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/shared/common/jwt"
	"github.com/mrakhaf/enqilo-store/shared/utils"
)

type (
	usecase struct {
		repository interfaces.Repository
		JwtAccess  *jwt.JWT
	}
)

func NewUsecase(repository interfaces.Repository, JwtAccess *jwt.JWT) interfaces.Usecase {
	return &usecase{
		repository: repository,
		JwtAccess:  JwtAccess,
	}
}

func (u *usecase) Login(ctx context.Context, phoneNumber, password string) (data dto.AuthResponse, err error) {

	token, err := u.JwtAccess.GenerateToken(phoneNumber)

	if err != nil {
		err = fmt.Errorf("failed to generate token: %s", err)
		return
	}

	data = dto.AuthResponse{
		PhoneNumber: phoneNumber,
		AccessToken: token,
	}

	return
}

func (u *usecase) Register(ctx context.Context, req request.Register) (data dto.AuthResponse, err error) {

	utils.IsValidPhoneNumber(req.PhoneNumber)

	//save db
	idStaff, err := u.repository.SaveUserAccount(req)

	if err != nil {
		err = fmt.Errorf("failed to save user account: %s", err)
		return
	}

	//generate token
	token, err := u.JwtAccess.GenerateToken(req.PhoneNumber)

	if err != nil {
		err = fmt.Errorf("failed to generate token: %s", err)
		return
	}

	data = dto.AuthResponse{
		PhoneNumber: req.PhoneNumber,
		AccessToken: token,
		Id:          idStaff,
		Name:        req.Name,
	}

	return
}

func (u *usecase) CheckIsUserExist(ctx context.Context, phoneNumber string) (isUserExist bool, data entity.User, err error) {
	data, err = u.repository.GetDataAccount(phoneNumber)

	if err != nil && err.Error() == "sql: no rows in result set" {
		err = nil
		isUserExist = false
		return
	}

	if err != nil {
		err = fmt.Errorf("failed to get data account: %s", err)
		return
	}

	isUserExist = true

	if err != nil {
		err = fmt.Errorf("failed to check is email exist: %s", err)
		return
	}

	return
}

func (u *usecase) CheckUserPassword(ctx context.Context, password string, data entity.User) (isPasswordCorrect bool) {

	err := utils.CheckPasswordHash(password, data.Password)

	if err != nil {
		isPasswordCorrect = false
		return
	}

	isPasswordCorrect = true

	return
}
