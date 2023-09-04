package user_v1

import (
	"context"

	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
)

func (i *Implementation) AddUser(ctx context.Context, req *desc.AddUserRequest) (*desc.AddUserResponse, error) {
	id, err := i.userService.AddUser(ctx, req.GetUserName())
	if err != nil {
		return nil, err
	}

	return &desc.AddUserResponse{
		Id: id,
	}, nil
}
