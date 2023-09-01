package user

import (
	"context"

	"github.com/nikitads9/segment-service-api/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ModifySegments(ctx context.Context, mod *model.ModifySegmentInfo) error {
	if mod == nil {
		return status.Error(codes.InvalidArgument, "received invalid request")
	}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		for _, elem := range mod.SlugsToAdd {
			errTx = s.userRepository.AddToSegment(ctx, elem, mod.UserId)
			if errTx != nil {
				return errTx
			}
		}

		for _, elem := range mod.SlugsToRemove {
			errTx = s.userRepository.RemoveFromSegment(ctx, elem, mod.UserId)
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
