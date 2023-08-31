package user

import "context"

func (s *Service) GetSegments(ctx context.Context, id int64) ([]string, error) {
	return s.userRepository.GetSegments(ctx, id)
}
