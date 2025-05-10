package data

import (
	v1 "agents/api/commission/service/v1"
	"agents/app/authn/service/internal/biz"
	"context"
)

var _ biz.CommissionRepo = (*commissionRepo)(nil)

type commissionRepo struct {
	data *Data
}

func NewCommissionRepo(data *Data) biz.CommissionRepo {
	return &commissionRepo{
		data: data,
	}
}

// InitUserCommission implements biz.CommissionRepo.
func (c *commissionRepo) InitUserCommission(ctx context.Context, userId string) error {
	_, err := c.data.cc.InitUserCommission(ctx, &v1.InitUserCommissionReq{
		UserId: userId,
	})
	return err
}
