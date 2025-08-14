package user

import (
	"context"

	"github.com/atimot/app/domain/user"
)

type UnregisterUsecase struct {
	userRepository user.UserRepository
}

func NewUnregisterUsecase(
	userRepository user.UserRepository,
) *UnregisterUsecase {
	return &UnregisterUsecase{
		userRepository: userRepository,
	}
}

func (uu *UnregisterUsecase) Run(ctx context.Context, input UnregisterUsecaseInputDTO) error {
	u, err := uu.userRepository.FindById(ctx, input.ID)
	if err != nil || u == nil {
		return err
	}
	if err := uu.userRepository.Delete(ctx, u); err != nil {
		return err
	}
	return nil
}
