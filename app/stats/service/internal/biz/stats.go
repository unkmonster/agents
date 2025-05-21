package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type RegisterEvent struct {
	Domain string
	UserId string
}

type RechargeEvent struct {
	Domain  string
	UserId  string
	Amount  int64
	Product string
}

type User struct {
	Id        string    `db:"id"`
	CreatedAT time.Time `db:"created_at"`
}

// 事件类型：Register, Recharge
type StatsRepo interface {
	AddRegister(ctx context.Context, domain string) error
	AddRecharge(ctx context.Context, domain string, amount int64) error
}

type DomainRepo interface {
	Get(ctx context.Context, domain string) (*User, error)
}

type CommissionRepo interface {
	// IncUserCommission 懒得改了，增加跟这笔订单相关的所有代理的余额
	IncUserCommission(ctx context.Context, domain string, amount int64) error
	IncUserRegistrationCount(ctx context.Context, userId string) error
}

type StatsUseCase struct {
	stats      StatsRepo
	log        *log.Helper
	domain     DomainRepo
	commission CommissionRepo
}

func NewStatsUseCase(repo StatsRepo, logger log.Logger, domain DomainRepo, commission CommissionRepo) *StatsUseCase {
	return &StatsUseCase{
		stats:      repo,
		log:        log.NewHelper(logger),
		domain:     domain,
		commission: commission,
	}
}

func (uc *StatsUseCase) HandleRegisterEvent(ctx context.Context, event *RegisterEvent) error {
	if err := uc.stats.AddRegister(ctx, event.Domain); err != nil {
		return err
	}

	user, err := uc.domain.Get(ctx, event.Domain)
	if err != nil {
		return err
	}

	if err := uc.commission.IncUserRegistrationCount(ctx, user.Id); err != nil {
		return err
	}

	return nil
}

func (uc *StatsUseCase) HandleRechargeEvent(ctx context.Context, event *RechargeEvent) error {
	if err := uc.stats.AddRecharge(ctx, event.Domain, event.Amount); err != nil {
		return err
	}

	if err := uc.commission.IncUserCommission(ctx, event.Domain, event.Amount); err != nil {
		return err
	}

	return nil
}
