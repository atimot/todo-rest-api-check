package user

import (
	"context"

	"github.com/atimot/app/domain/errors"
)

type userDomainService struct {
	userRepository UserRepository
}

func NewUserDomainService(userRepository UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}

func (uds userDomainService) IsExists(ctx context.Context, email Email) (bool, error) {
	user, err := uds.userRepository.FindByEmail(ctx, email)
	if err != nil && errors.Is(err, errors.ErrNotFoundUser) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
