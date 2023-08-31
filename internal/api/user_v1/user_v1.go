package user_v1

import (
	"github.com/nikitads9/segment-service-api/internal/service/user"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
)

type Implementation struct {
	desc.UnimplementedUserV1ServiceServer
	userService *user.Service
}

func NewImplementation(userService *user.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedUserV1ServiceServer{},
		userService,
	}
}
