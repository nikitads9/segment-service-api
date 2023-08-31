package user

import (
	"context"
	"time"
)

func (s *Service) SetExpireTime(ctx context.Context, userId int64, slug string, expiration time.Time) error {
	return s.userRepository.SetExpireTime(ctx, userId, slug, expiration)
}
