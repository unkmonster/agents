package data

import (
	v1 "agents/api/user/service/v1"
	"agents/app/commission/service/internal/biz"
	"context"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

// GetUser implements biz.UserRepo.
func (u *userRepo) GetUser(ctx context.Context, userId string) (*biz.User, error) {
	reply, err := u.data.uc.GetUser(ctx, &v1.GetUserRequest{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:           reply.Id,
		Username:     reply.Username,
		Nickname:     reply.Nickname,
		ParentId:     reply.ParentId,
		Level:        reply.Level,
		SharePercent: reply.SharePercent,
	}, nil
}

// GetUserByDomain implements biz.UserRepo.
func (u *userRepo) GetUserByDomain(ctx context.Context, domain string) (*biz.User, error) {
	reply, err := u.data.uc.GetUserByDomain(ctx, &v1.GetUserByDomainRequest{
		Domain: domain,
	})

	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:           reply.Id,
		Username:     reply.Username,
		Nickname:     reply.Nickname,
		ParentId:     reply.ParentId,
		Level:        reply.Level,
		SharePercent: reply.SharePercent,
	}, nil
}
