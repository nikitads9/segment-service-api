package user_v1

import (
	"context"

	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SetExpireTime(ctx context.Context, req *desc.SetExpireTimeRequest) (*emptypb.Empty, error) {
	err := i.userService.SetExpireTime(ctx, req.GetId(), req.GetSlug(), req.GetExpirationTime().AsTime())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
