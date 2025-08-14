package user

import (
	"context"
	"testing"

	"github.com/atimot/app/domain/errors"
	"github.com/atimot/app/domain/user"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func TestUser_UpdateProfileUsecase_Run(t *testing.T) {
	tests := []struct {
		name    string
		mockFn  func(mr *user.MockUserRepository)
		input   UpdateProfileUsecaseInputDTO
		want    *UpdateProfileUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系:プロフィールを編集できる",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(user.ReconstructUser("1", "user@test.com", "testuser", ""), nil)
				mr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: UpdateProfileUsecaseInputDTO{
				ID:    "1",
				Email: "updated@test.com",
				Name:  "updateduser",
			},
			want: &UpdateProfileUsecaseOutputDTO{
				ID:    "1",
				Email: "updated@test.com",
				Name:  "updateduser",
			},
			wantErr: false,
		},
		{
			name: "正常系:空のフィールドは更新されない",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(user.ReconstructUser("1", "user@test.com", "testuser", ""), nil)
				mr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: UpdateProfileUsecaseInputDTO{
				ID:   "1",
				Name: "updateduser",
			},
			want: &UpdateProfileUsecaseOutputDTO{
				ID:    "1",
				Email: "user@test.com",
				Name:  "updateduser",
			},
			wantErr: false,
		},
		{
			name: "準正常系:存在しないユーザーは編集できない",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(nil, errors.ErrNotFoundUser)
			},
			input: UpdateProfileUsecaseInputDTO{
				ID:    "0",
				Email: "noexsitent@test.com",
				Name:  "noexsitentuser",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUserRepository := user.NewMockUserRepository(ctrl)
			tt.mockFn(mockUserRepository)

			updateProfileUsecase := NewUpdateProfileUsecase(mockUserRepository)
			ctx := context.Background()
			got, err := updateProfileUsecase.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfileUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("UpdateProfileUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
