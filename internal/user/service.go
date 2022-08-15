package user

import (
	"context"
	"go-rest-api-pet/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	//TODO for the next one
	return
}
