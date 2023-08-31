package user

import "context"

func (s *Service) AddUser(ctx context.Context, userName string) (int64, error) {
	return s.userRepository.AddUser(ctx, userName)
}
