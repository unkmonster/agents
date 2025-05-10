package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	commissionv1 "agents/api/commission/service/v1"
)

type Commission struct {
	Id                string `db:"id"`
	UserId            string `db:"user_id"`
	TodayCommission   int32  `db:"today_commission"`
	TotalCommission   int32  `db:"total_commission"`
	SettledCommission int32  `db:"settled_commission"`
}

type CommissionRepo interface {
	InitUserCommission(ctx context.Context, userId string) error
	IncUserCommission(ctx context.Context, userId string, amount int32) error
	IncUserSettledCommission(ctx context.Context, userId string, amount int32) error
	GetUserCommission(ctx context.Context, userId string) (*Commission, error)
	ListCommission(ctx context.Context) ([]*Commission, error)
	// ListCommissionByParent 列出直接子用户的佣金
	ListCommissionByParent(ctx context.Context, parentId string) ([]*Commission, error)
}

type CommissionUseCase struct {
	commission CommissionRepo
	log        *log.Helper
	user       UserRepo
}

func NewCommissionUseCase(commission CommissionRepo, logger log.Logger, user UserRepo) *CommissionUseCase {
	return &CommissionUseCase{
		commission: commission,
		log:        log.NewHelper(logger),
		user:       user,
	}
}

func (uc *CommissionUseCase) CalcOrderCommission(ctx context.Context, req *commissionv1.HandleOrderCommissionRequest) error {
	user, err := uc.user.GetUserByDomain(ctx, req.Domain)
	if err != nil {
		return err
	}

	stk := []*User{}
	stk = append(stk, user)

	for user.Level > 1 {
		user, err = uc.user.GetUser(ctx, *user.ParentId)
		if err != nil {
			return err
		}
		stk = append(stk, user)
	}

	// 计算佣金
	commissions := []float32{}

	for i := len(stk) - 1; i >= 0; i-- {
		if len(commissions) == 0 {
			commissions = append(commissions, float32(req.Amount)*stk[i].SharePercent)
		} else {
			commissions = append(commissions, float32(commissions[len(commissions)-1])*stk[i].SharePercent)
		}
	}

	n := len(stk)
	for i := range n {
		user := stk[n-i-1]
		amount := int32(0)

		if i+1 < n {
			amount = int32(commissions[i]) - int32(commissions[i+1])
		} else {
			amount = int32(commissions[i])
		}

		if err := uc.commission.IncUserCommission(ctx, user.Id, amount); err != nil {
			return err
		}
	}

	return nil
}

func (uc *CommissionUseCase) InitUserCommission(ctx context.Context, userId string) error {
	return uc.commission.InitUserCommission(ctx, userId)
}

func (uc *CommissionUseCase) InUserCommission(ctx context.Context, userId string, amount int32) error {
	return uc.commission.IncUserCommission(ctx, userId, amount)
}

func (uc *CommissionUseCase) IncUserSettledCommission(ctx context.Context, userId string, amount int32) error {
	return uc.commission.IncUserSettledCommission(ctx, userId, amount)
}

func (uc *CommissionUseCase) GetUserCommission(ctx context.Context, userId string) (*Commission, error) {
	return uc.commission.GetUserCommission(ctx, userId)
}

func (uc *CommissionUseCase) ListCommission(ctx context.Context) ([]*Commission, error) {
	return uc.commission.ListCommission(ctx)
}

func (uc *CommissionUseCase) ListCommissionByParent(ctx context.Context, parentId string) ([]*Commission, error) {
	return uc.commission.ListCommissionByParent(ctx, parentId)
}
