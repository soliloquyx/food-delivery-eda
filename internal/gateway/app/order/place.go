package order

import (
	"context"
)

func (s *service) Place(ctx context.Context, in PlaceInput) (PlaceResult, error) {
	result, err := s.client.Place(ctx, in)
	if err != nil {
		return PlaceResult{}, nil
	}

	return result, nil
}
