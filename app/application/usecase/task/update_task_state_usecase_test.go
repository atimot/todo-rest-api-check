package task

import (
	"context"
	"testing"

	"github.com/atimot/app/domain/errors"
	"github.com/atimot/app/domain/task"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func TestTask_UpdateStateUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		mockFn  func(mr *task.MockTaskRepository)
		input   UpdateTaskStateUsecaseInputDTO
		errType error
		want    *UpdateTaskStateUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					task.ReconstructTask("id", "user_id", "this is content", 0),
					nil,
				)
				mr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: UpdateTaskStateUsecaseInputDTO{
				ID:     "id",
				UserId: "user_id",
				State:  "done",
			},
			want: &UpdateTaskStateUsecaseOutputDTO{
				ID:      "id",
				UserId:  "user_id",
				Content: "this is content",
				State:   "done",
			},
			wantErr: false,
		},
		{
			name: "準正常系:指定したstateが不正",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					task.ReconstructTask("id", "user_id", "this is content", 0),
					nil,
				)
			},
			errType: errors.ErrInvalidTaskState,
			input: UpdateTaskStateUsecaseInputDTO{
				ID:     "id",
				UserId: "user_id",
				State:  "invalid",
			},
			wantErr: true,
		},
		{
			name: "準正常系:指定したidを持つタスクが存在しない",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					nil,
					errors.ErrNotFoundTask,
				)
			},
			errType: errors.ErrNotFoundTask,
			input: UpdateTaskStateUsecaseInputDTO{
				ID:     "id",
				UserId: "user_id",
				State:  "doing",
			},
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
			input: UpdateTaskStateUsecaseInputDTO{
				ID:     "id",
				UserId: "other_user_id",
				State:  "doing",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockTaskRepository := task.NewMockTaskRepository(ctrl)
			tt.mockFn(mockTaskRepository)

			sut := NewUpdateTaskStateUsecase(mockTaskRepository)
			ctx := context.Background()
			got, err := sut.Run(ctx, tt.input)

			if tt.errType != nil && !errors.Is(err, tt.errType) {
				t.Errorf("updateTaskState.Run = error:%v,want errYType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("updateTaskState.Run = error:%v,wantErr:%v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("updateTaskState.Run() -got,+want :%v ", diff)
				return
			}
		})
	}
}
