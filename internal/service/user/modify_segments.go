package user

import "context"

func (s *Service) ModifySegments(ctx context.Context, slugsToAdd []string, slugsToRemove []string, id int64) error {
	return s.userRepository.ModifySegments(ctx, slugsToAdd, slugsToRemove, id)
}
