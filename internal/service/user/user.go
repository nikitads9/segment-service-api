package user

import (
	"github.com/nikitads9/segment-service-api/internal/repository/user"
)

type Service struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}
