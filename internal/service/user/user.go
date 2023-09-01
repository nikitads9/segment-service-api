package user

import (
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/user"
)

type Service struct {
	userRepository user.Repository
	txManager      db.TxManager
}

func NewUserService(userRepository user.Repository, txManager db.TxManager) *Service {
	return &Service{
		userRepository: userRepository,
		txManager:      txManager,
	}
}

func NewMockUserService(deps ...interface{}) *Service {
	is := Service{}
	for _, val := range deps {
		switch s := val.(type) {
		case user.Repository:
			is.userRepository = s
		}
	}
	return &is
}
