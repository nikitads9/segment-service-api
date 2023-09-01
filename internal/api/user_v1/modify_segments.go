package user_v1

import (
	"context"

	"github.com/nikitads9/segment-service-api/internal/convert"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) ModifySegments(ctx context.Context, req *desc.ModifySegmentsRequest) (*emptypb.Empty, error) {
	err := i.userService.ModifySegments(ctx, convert.ToModifySegmentInfo(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
