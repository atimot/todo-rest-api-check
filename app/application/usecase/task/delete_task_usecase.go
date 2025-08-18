package task

import (
	"context"

	"github.com/atimot/app/domain/task"
)

type DeleteTaskUsecase struct {
	taskRepository task.TaskRepository
}

func NewDeleteTaskUsecase(taskRepository task.TaskRepository) *DeleteTaskUsecase {
	return &DeleteTaskUsecase{
		taskRepository: taskRepository,
	}
}

func (dtu *DeleteTaskUsecase) Run(ctx context.Context, input DeleteTaskUsecaseInputDTO) error {
	t, err := dtu.taskRepository.FindById(ctx, input.ID)
	if err != nil {
		return err
	}

	// ユーザーがタスクを操作できるかどうかを確認
	if err := t.IsOperableBy(input.UserId); err != nil {
		return err
	}

	if err := dtu.taskRepository.Delete(ctx, t); err != nil {
		return err
	}
	return nil
}
