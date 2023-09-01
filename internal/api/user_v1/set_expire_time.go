package user_v1

import (
	"context"

	"github.com/nikitads9/segment-service-api/internal/convert"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SetExpireTime(ctx context.Context, req *desc.SetExpireTimeRequest) (*emptypb.Empty, error) {
	err := i.userService.SetExpireTime(ctx, convert.ToSetExpireTimeInfo(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
