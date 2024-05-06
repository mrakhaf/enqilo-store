package interfaces

import (
	"context"

	"github.com/mrakhaf/enqilo-store/models/dto"
	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
)

type Usecase interface {
	Login(ctx context.Context, email, password string) (data dto.AuthResponse, err error)
	Register(ctx context.Context, req request.Register) (data dto.AuthResponse, err error)
	CheckIsUserExist(ctx context.Context, phoneNumber string) (isUserExist bool, data entity.User, err error)
	CheckUserPassword(ctx context.Context, password string, data entity.User) (isPasswordCorrect bool)
}
