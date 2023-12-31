package user_v1

import (
	"context"

	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) RemoveUser(ctx context.Context, req *desc.RemoveUsertRequest) (*emptypb.Empty, error) {
	err := i.userService.RemoveUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
