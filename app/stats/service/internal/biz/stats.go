package biz

import (
	"context"

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

// 事件类型：Register, Recharge
type StatsRepo interface {
	AddRegister(ctx context.Context, domain string) error
	AddRecharge(ctx context.Context, domain string, amount int64) error
}

type StatsUseCase struct {
	repo StatsRepo
	log  *log.Helper
}

func NewStatsUseCase(repo StatsRepo, logger log.Logger) *StatsUseCase {
	return &StatsUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *StatsUseCase) HandleRegisterEvent(ctx context.Context, event *RegisterEvent) error {
	return uc.repo.AddRegister(ctx, event.Domain)
}

func (uc *StatsUseCase) HandleRechargeEvent(ctx context.Context, event *RechargeEvent) error {
	return uc.repo.AddRecharge(ctx, event.Domain, event.Amount)
}
