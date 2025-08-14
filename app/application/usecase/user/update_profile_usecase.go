package user

import (
	"context"

	"github.com/atimot/app/domain/user"
)

type UpdateProfileUsecase struct {
	userRepository user.UserRepository
}

func NewUpdateProfileUsecase(userRepository user.UserRepository) *UpdateProfileUsecase {
	return &UpdateProfileUsecase{
		userRepository: userRepository,
	}
}

func (epu *UpdateProfileUsecase) Run(ctx context.Context, input UpdateProfileUsecaseInputDTO) (*UpdateProfileUsecaseOutputDTO, error) {
	// 存在しているユーザーしか編集できない
	u, err := epu.userRepository.FindById(ctx, input.ID)
	if err != nil || u == nil {
		return nil, err
	}
	if input.Name == "" {
		input.Name = u.GetName()
	}
	if input.Email == "" {
		input.Email = u.GetEmail().Value()
	}

	updatedUser, err := u.UpdateUser(
		input.Email,
		input.Name,
	)
	if err != nil {
		return nil, err
	}
	if err := epu.userRepository.Update(ctx, updatedUser); err != nil {
		return nil, err
	}

	return &UpdateProfileUsecaseOutputDTO{
		ID:    updatedUser.GetID(),
		Email: updatedUser.GetEmail().Value(),
		Name:  updatedUser.GetName(),
	}, nil
}
