package user

import "context"

func (s *Service) GetSegments(ctx context.Context, id int64) ([]string, error) {
	/* 	var res []string
	   	segments, err := s.userRepository.GetSegments(ctx, id)
	   	if err != nil {
	   		return nil, err
	   	}
	   	for _, elem := range segments {
	   		res = append(res, elem.Slug)
	   	}
	   	return res, nil */
	return s.userRepository.GetSegments(ctx, id)
}
