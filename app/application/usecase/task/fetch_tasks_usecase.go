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

func (ftu *FetchTasksUsecase) Run(ctx context.Context) ([]*FetchTaskUsecaseOutputDTO, error) {
	dtos, err := ftu.taskQueryService.FetchAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	outputs := make([]*FetchTaskUsecaseOutputDTO, 0, len(dtos))
	for _, dto := range dtos {
		outputs = append(outputs, &FetchTaskUsecaseOutputDTO{
			ID:       dto.ID,
			UserId:   dto.UserId,
			UserName: dto.UserName,
			Content:  dto.Content,
			State:    dto.State,
		})
	}
	return outputs, nil
}
