package user

import "context"

func (s *Service) RemoveUser(ctx context.Context, id int64) error {
	return s.userRepository.RemoveUser(ctx, id)
}
