package task

import "context"

type FetchTasksUsecase struct {
	taskQueryService TaskQueryService
}

func NewFetchTasksUsecase(taskQueryService TaskQueryService) *FetchTasksUsecase {
	return &FetchTasksUsecase{
		taskQueryService: taskQueryService,
	}
}

func (ftu *FetchTasksUsecase) Run(ctx context.Context, input FetchTaskUsecaseInputDTO) (*FetchTaskUsecaseOutputDTO, error) {
	dto, err := ftu.taskQueryService.FetchTaskById(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &FetchTaskUsecaseOutputDTO{
		ID:       dto.ID,
		UserId:   dto.UserId,
		UserName: dto.UserName,
		Content:  dto.Content,
		State:    dto.State,
	}, nil
}
