package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	commissionv1 "agents/api/commission/service/v1"
)

const (
	CommissionTypeDirect   = "direct"
	CommissionTypeIndirect = "indirect"
)

type Commission struct {
	Id                     string `db:"id"`
	UserId                 string `db:"user_id"`
	TodayCommission        int32  `db:"today_commission"`
	TotalCommission        int32  `db:"total_commission"`
	SettledCommission      int32  `db:"settled_commission"`
	TodayRegistrationCount int64  `db:"today_registration_count"`
	TotalRegistrationCount int64  `db:"total_registration_count"`
}

type CommissionRepo interface {
	InitUserCommission(ctx context.Context, userId string) error
	IncUserCommission(ctx context.Context, userId string, amount int32) error
	IncUserSettledCommission(ctx context.Context, userId string, amount int32) error
	GetUserCommission(ctx context.Context, userId string) (*Commission, error)
	ListCommission(ctx context.Context) ([]*Commission, error)
	// ListCommissionByParent 列出直接子用户的佣金
	ListCommissionByParent(ctx context.Context, parentId string) ([]*Commission, error)
	IncUserRegistrationCount(ctx context.Context, userId string) error

	// ----------- daily -------------------
	IncUserDirectCommission(ctx context.Context, userId string, amount int64) error
	IncUserIndirectCommission(ctx context.Context, userId string, amount int64) error
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

	for _, u := range stk {
		uc.log.Debugf("user: %+v", u)
	}

	// 计算佣金，未扣除子代理的分成
	commissions := []float32{float32(req.Amount)} // index equal to agent level

	for i := len(stk) - 1; i >= 0; i-- {
		commissions = append(commissions, commissions[len(commissions)-1]*stk[i].SharePercent)
	}
	uc.log.Debugf("commissions: %+v", commissions)

	// 子代理的分成比例来自其父代理
	n := len(stk)
	for i := range n {
		direct := false

		user := stk[n-1-i]
		amount := int32(commissions[user.Level])
		if int(user.Level+1) < len(commissions) {
			amount -= int32(commissions[user.Level+1])
		} else {
			direct = true
		}

		uc.log.Debugf("user: %+v, amount: %d", user, amount)

		var err error
		if direct {
			err = uc.commission.IncUserDirectCommission(ctx, user.Id, int64(amount))
		} else {
			err = uc.commission.IncUserIndirectCommission(ctx, user.Id, int64(amount))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (uc *CommissionUseCase) InitUserCommission(ctx context.Context, userId string) error {
	return uc.commission.InitUserCommission(ctx, userId)
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

func (uc *CommissionUseCase) IncUserRegistrationCount(ctx context.Context, userId string) error {
	return uc.commission.IncUserRegistrationCount(ctx, userId)
}
