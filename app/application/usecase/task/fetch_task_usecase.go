package task

import "context"

type FetchTaskUsecase struct {
	taskQueryService TaskQueryService
}

func NewFetchTaskUsecase(taskQueryService TaskQueryService) *FetchTaskUsecase {
	return &FetchTaskUsecase{
		taskQueryService: taskQueryService,
	}
}

func (ftu *FetchTaskUsecase) Run(ctx context.Context, input FetchTaskUsecaseInputDTO) (*FetchTaskUsecaseOutputDTO, error) {
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
