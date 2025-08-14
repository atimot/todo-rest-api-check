package user

import (
	"context"
	"testing"

	"github.com/atimot/app/domain/errors"
	"github.com/atimot/app/domain/user"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"
)

func TestUser_RegisterUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		mockFn  func(mr *user.MockUserRepository, mds *user.MockUserDomainService)
		input   RegisterUsecaseInputDTO
		errType error
		want    *RegisterUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系:ユーザーを登録できる",
			mockFn: func(mr *user.MockUserRepository, mds *user.MockUserDomainService) {
				mds.EXPECT().IsExists(gomock.Any(), gomock.Any()).Return(false, nil)
				mr.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: RegisterUsecaseInputDTO{
				Name:     "user",
				Email:    "email@test.com",
				Password: "password",
			},
			want: &RegisterUsecaseOutputDTO{
				Name:  "user",
				Email: "email@test.com",
			},
			wantErr: false,
		},
		{
			name: "準正常系:重複した登録は阻止する",
			mockFn: func(mr *user.MockUserRepository, mds *user.MockUserDomainService) {
				// emailで検索するとDBに存在していたケース
				mds.EXPECT().IsExists(gomock.Any(), gomock.Any()).Return(true, nil)
			},
			input: RegisterUsecaseInputDTO{
				Name:     "user",
				Email:    "email@test.com",
				Password: "password",
			},
			errType: errors.ErrAlreadyRegisterd,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUserRepository := user.NewMockUserRepository(ctrl)
			mockUserDomainService := user.NewMockUserDomainService(ctrl)
			tt.mockFn(mockUserRepository, mockUserDomainService)

			registerUsecase := NewRegisterUsecase(mockUserRepository, mockUserDomainService)
			ctx := context.Background()
			got, err := registerUsecase.Run(ctx, tt.input)

			if tt.errType != nil && !errors.Is(err, tt.errType) {
				t.Errorf("registerUsecase.Run = error:%v,want errYType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("registerUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(RegisterUsecaseOutputDTO{}, "ID")); diff != "" {
				t.Errorf("registerUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
