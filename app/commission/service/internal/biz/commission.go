package biz

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	commissionv1 "agents/api/commission/service/v1"
	"agents/pkg/paging"
)

const (
	CommissionTypeDirect   = "direct"
	CommissionTypeIndirect = "indirect"
)

type TotalCommission struct {
	Id                     string `db:"id"`
	UserId                 string `db:"user_id"`
	TodayCommission        int32  `db:"today_commission"`
	TotalCommission        int32  `db:"total_commission"`
	SettledCommission      int32  `db:"settled_commission"`
	TodayRegistrationCount int64  `db:"today_registration_count"`
	TotalRegistrationCount int64  `db:"total_registration_count"`
}

type DailyCommission struct {
	Date                      time.Time `db:"date"`
	UserId                    string    `db:"user_id"`
	IndirectRechargeAmount    int64     `db:"indirect_recharge_amount"`
	DirectRechargeAmount      int64     `db:"direct_recharge_amount"`
	IndirectRegistrationCount int64     `db:"indirect_registration_count"`
	DirectRegistrationCount   int64     `db:"direct_registration_count"`
	UpdatedAt                 time.Time `db:"updated_at"`
}

type CommissionRepo interface {
	IncUserSettledCommission(ctx context.Context, userId string, amount int32) error
	GetUserTotalCommission(ctx context.Context, userId string) (*TotalCommission, error)
	// ListCommission 列出系统内所有用户的累计佣金
	ListCommission(ctx context.Context) ([]*TotalCommission, error)
	// ListTotalCommissionByParent 列出直接子用户的累计佣金
	ListTotalCommissionByParent(ctx context.Context, parentId string) ([]*TotalCommission, error)
	IncUserRegistrationCount(ctx context.Context, userId string) error

	// ----------- daily -------------------
	IncUserDirectCommission(ctx context.Context, userId string, amount int64) error
	IncUserIndirectCommission(ctx context.Context, userId string, amount int64) error
	IncUserDirectRegistrationCount(ctx context.Context, userId string) error
	IncUserIndirectRegistrationCount(ctx context.Context, userId string) error

	GetUserCommissionByDate(ctx context.Context, userId string, date time.Time) (*DailyCommission, error)
	ListCommissionByUser(ctx context.Context, userId string, paging *paging.Paging) ([]*DailyCommission, error)
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

func (uc *CommissionUseCase) GetUserTotalCommission(ctx context.Context, userId string) (*TotalCommission, error) {
	return uc.commission.GetUserTotalCommission(ctx, userId)
}

func (uc *CommissionUseCase) ListTotalCommission(ctx context.Context) ([]*TotalCommission, error) {
	return uc.commission.ListCommission(ctx)
}

func (uc *CommissionUseCase) ListTotalCommissionByParent(ctx context.Context, parentId string) ([]*TotalCommission, error) {
	return uc.commission.ListTotalCommissionByParent(ctx, parentId)
}

// IncChainRegistrationCountByDirectUser 增加从 user_id 到他的顶级代理每个代理的注册量
func (uc *CommissionUseCase) IncChainRegistrationCountByDirectUser(ctx context.Context, userId string) error {
	user, err := uc.user.GetUser(ctx, userId)
	if err != nil {
		return err
	}

	users := []*User{user}

	for user.Level > 1 {
		var err error
		user, err = uc.user.GetUser(ctx, *user.ParentId)
		if err != nil {
			return err
		}
		users = append(users, user)
	}

	for i, user := range users {
		if i == 0 {
			err = uc.commission.IncUserDirectRegistrationCount(ctx, user.Id)
		} else {
			err = uc.commission.IncUserIndirectRegistrationCount(ctx, user.Id)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *CommissionUseCase) ListCommissionByUser(ctx context.Context, req *commissionv1.ListCommissionByUserReq) ([]*DailyCommission, error) {
	if req.Date != "" {
		date, err := time.Parse(time.DateOnly, req.Date)
		if err != nil {
			return nil, err
		}
		comm, err := uc.commission.GetUserCommissionByDate(ctx, req.UserId, date)
		if err == sql.ErrNoRows {
			return []*DailyCommission{{
				UserId: req.UserId,
				Date:   date,
			}}, nil
		} else if err != nil {
			return nil, err
		}
		return []*DailyCommission{comm}, nil
	}

	return uc.commission.ListCommissionByUser(ctx, req.UserId, &paging.Paging{
		Offset:  int64(req.Offset),
		Limit:   int64(req.Limit),
		OrderBy: req.OrderBy,
		Sort:    req.Sort,
	})
}
