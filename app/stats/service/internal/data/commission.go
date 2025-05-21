package data

import (
	v1 "agents/api/commission/service/v1"
	"agents/app/stats/service/internal/biz"
	"context"
)

type commissionRepo struct {
	data *Data
}

// IncUserCommission implements biz.CommissionRepo.
func (c *commissionRepo) IncUserCommission(ctx context.Context, domain string, amount int64) error {
	_, err := c.data.cc.HandleOrderCommission(ctx, &v1.HandleOrderCommissionRequest{
		Domain: domain,
		Amount: int32(amount),
	})
	return err
}

// IncUserRegistrationCount implements biz.CommissionRepo.
func (c *commissionRepo) IncUserRegistrationCount(ctx context.Context, userId string) error {
	_, err := c.data.cc.IncUserRegistrationCount(ctx, &v1.IncUserRegistrationCountReq{
		UserId: userId,
	})
	return err
}

func NewCommissionRepo(data *Data) biz.CommissionRepo {
	return &commissionRepo{
		data: data,
	}
}
