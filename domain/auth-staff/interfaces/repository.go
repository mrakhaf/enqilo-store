package interfaces

import (
	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
)

type Repository interface {
	SaveUserAccount(data request.Register) (id string, err error)
	GetDataAccount(email string) (data entity.User, err error)
}
