package data

import (
	userv1 "agents/api/user/service/v1"
	"agents/app/authn/service/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create implements biz.UserRepo.
func (u *userRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	reply, err := u.data.uc.CreateUser(ctx, &userv1.CreateUserRequest{
		Username:     &user.Username,
		Nickname:     user.Nickname,
		ParentId:     user.ParentId,
		Level:        &user.Level,
		SharePercent: &user.SharePercent,
	})

	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:           *reply.Id,
		Username:     *reply.Username,
		Nickname:     reply.Nickname,
		ParentId:     reply.ParentId,
		Level:        *reply.Level,
		SharePercent: *reply.SharePercent,
	}, nil
}

// GetByUsername implements biz.UserRepo.
func (u *userRepo) GetByUsername(ctx context.Context, username string) (*biz.User, error) {
	reply, err := u.data.uc.GetUserByUsername(ctx, &userv1.GetUserByUsernameRequest{
		Username: &username,
	})

	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:           *reply.Id,
		Username:     *reply.Username,
		Nickname:     reply.Nickname,
		ParentId:     reply.ParentId,
		Level:        *reply.Level,
		SharePercent: *reply.SharePercent,
	}, nil
}
