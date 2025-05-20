package service

import (
	"context"
	"strings"

	pb "agents/api/stats/service/v1"
	"agents/app/stats/service/internal/biz"
)

const (
	EventUserRegister = "user_register"
	EventUserRecharge = "user_recharge"
)

type StatsService struct {
	pb.UnimplementedStatsServer
	stats *biz.StatsUseCase
}

func NewStatsService(stats *biz.StatsUseCase) *StatsService {
	return &StatsService{
		stats: stats,
	}
}

func (s *StatsService) CreateEvent(ctx context.Context, req *pb.CreateEventReq) (*pb.CreateEventReply, error) {
	switch strings.ToLower(req.Type) {
	case EventUserRecharge:
		err := s.stats.HandleRechargeEvent(ctx, &biz.RechargeEvent{
			Domain:  req.GetRecharge().Domain,
			UserId:  req.GetRecharge().UserId,
			Amount:  req.GetRecharge().Amount,
			Product: req.GetRecharge().Product,
		})
		return nil, err
	case EventUserRegister:
		reg := req.GetRegister()
		err := s.stats.HandleRegisterEvent(ctx, &biz.RegisterEvent{
			Domain: reg.Domain,
			UserId: reg.UserId,
		})
		return nil, err
	default:
		return nil, pb.ErrorUnknownEventType("event: %s", req.Type)
	}
}
