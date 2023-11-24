package home

import (
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
	"gorm.io/gorm"
)

type Handler struct {
	DbConnection *gorm.DB
}

func (h Handler) HandleHome(input server.Input) (o server.Output) {
	provider := user.Provider{
		DbConnection: h.DbConnection,
	}

	users, err := provider.GetAll()
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}
	return server.Output{
		Data: users,
		Code: 200,
	}
}
