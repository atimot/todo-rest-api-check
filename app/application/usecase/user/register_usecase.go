package user

import (
	"context"

	"github.com/atimot/app/domain/errors"
	"github.com/atimot/app/domain/user"
)

type RegisterUsecase struct {
	userRepository    user.UserRepository
	userDomainService user.UserDomainService
}

func NewRegisterUsecase(
	userRepository user.UserRepository,
	userDomainService user.UserDomainService,
) *RegisterUsecase {
	return &RegisterUsecase{
		userRepository:    userRepository,
		userDomainService: userDomainService,
	}
}

func (ru *RegisterUsecase) Run(ctx context.Context, input RegisterUsecaseInputDTO) (*RegisterUsecaseOutputDTO, error) {
	u, err := user.NewUser(
		input.Email,
		input.Name,
		input.Password,
	)
	if err != nil {
		return nil, err
	}
	ok, err := ru.userDomainService.IsExists(ctx, u.GetEmail())
	if err != nil {
		return nil, err
	}
	if ok {
		return nil, errors.ErrAlreadyRegisterd
	}
	if err := ru.userRepository.Save(ctx, u); err != nil {
		return nil, err
	}

	return &RegisterUsecaseOutputDTO{
		ID:    u.GetID(),
		Name:  u.GetName(),
		Email: u.GetEmail().Value(),
	}, nil
}
