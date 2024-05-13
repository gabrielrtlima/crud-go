package converter

import (
	"github.com/gabrielrtlima/crud-go/src/model"
	"github.com/gabrielrtlima/crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.Email, entity.Name, entity.Password, entity.Age)
	domain.SetID(entity.ID.Hex())
	return domain
}
