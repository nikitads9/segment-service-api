package user

import (
	"context"

	"github.com/nikitads9/segment-service-api/internal/model"
)

func (s *Service) SetExpireTime(ctx context.Context, mod *model.SetExpireTimeInfo) error {
	return s.userRepository.SetExpireTime(ctx, mod)
}
