package user

import "context"

func (s *Service) ModifySegments(ctx context.Context, slugsToAdd []string, slugsToRemove []string, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		for _, elem := range slugsToAdd {
			errTx = s.userRepository.AddToSegment(ctx, elem, id)
			if errTx != nil {
				return errTx
			}
		}

		for _, elem := range slugsToRemove {
			errTx = s.userRepository.RemoveFromSegment(ctx, elem, id)
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
