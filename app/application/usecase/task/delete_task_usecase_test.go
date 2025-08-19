package task

import (
	"context"
	"testing"

	"github.com/atimot/app/domain/errors"
	"github.com/atimot/app/domain/task"
	"go.uber.org/mock/gomock"
)

func TestTask_DeleteTaskUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		mockFn  func(mr *task.MockTaskRepository)
		input   DeleteTaskUsecaseInputDTO
		errType error
		wantErr bool
	}{
		{
			name: "正常系",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					task.ReconstructTask("id", "user_id", "this is content", 0),
					nil,
				)
				mr.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: DeleteTaskUsecaseInputDTO{
				ID:     "id",
				UserId: "user_id",
			},
			wantErr: false,
		},
		{
			name: "準正常系:指定したidを持つタスクが存在しない",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					nil,
					errors.ErrNotFoundTask,
				)
			},
			input: DeleteTaskUsecaseInputDTO{
				ID:     "id",
				UserId: "user_id",
			},
			errType: errors.ErrNotFoundTask,
			wantErr: true,
		},
		{
			name: "準正常系:異なるuser_idで操作するとエラーが返る",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					task.ReconstructTask("id", "user_id", "this is content", 0),
					nil,
				)
			},
			input: DeleteTaskUsecaseInputDTO{
				ID:     "id",
				UserId: "other_user_id",
			},
			errType: errors.ErrForbiddenTaskOperation,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockTaskRepository := task.NewMockTaskRepository(ctrl)
			tt.mockFn(mockTaskRepository)

			sut := NewDeleteTaskUsecase(mockTaskRepository)
			ctx := context.Background()
			err := sut.Run(ctx, tt.input)

			if tt.errType != nil && !errors.Is(err, tt.errType) {
				t.Errorf("deleteTaskUsecase.Run = error:%v,want errYType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteTaskUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
				return
			}
		})
	}
}
