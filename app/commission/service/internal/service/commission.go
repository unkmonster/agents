package service

import (
	"context"

	pb "agents/api/commission/service/v1"
	"agents/app/commission/service/internal/biz"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommissionService struct {
	pb.UnimplementedCommissionServer
	comm *biz.CommissionUseCase
}

func NewCommissionService(comm *biz.CommissionUseCase) *CommissionService {
	return &CommissionService{
		comm: comm,
	}
}

func (s *CommissionService) HandleOrderCommission(ctx context.Context, req *pb.HandleOrderCommissionRequest) (*pb.HandleOrderCommissionReply, error) {
	err := s.comm.CalcOrderCommission(ctx, req)
	return &pb.HandleOrderCommissionReply{}, err
}

func (s *CommissionService) IncChainRegistrationCountByDirectUser(ctx context.Context, req *pb.IncChainRegistrationCountByDirectUserReq) (*pb.IncChainRegistrationCountByDirectUserReply, error) {
	err := s.comm.IncChainRegistrationCountByDirectUser(ctx, req.UserId)
	return nil, err
}

func (s *CommissionService) GetUserTotalCommission(ctx context.Context, req *pb.GetUserTotalCommissionRequest) (*pb.GetUserTotalCommissionReply, error) {
	comm, err := s.comm.GetUserTotalCommission(ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	return &pb.GetUserTotalCommissionReply{
		Id:                     comm.Id,
		UserId:                 comm.UserId,
		TotalCommission:        comm.TotalCommission,
		TodayCommission:        comm.TodayCommission,
		SettledCommission:      comm.SettledCommission,
		TodayRegistrationCount: int32(comm.TodayRegistrationCount),
		TotalRegistrationCount: int32(comm.TotalRegistrationCount),
	}, nil
}

func (s *CommissionService) ListTotalCommission(ctx context.Context, req *pb.ListTotalCommissionRequest) (*pb.ListTotalCommissionReply, error) {
	comms, err := s.comm.ListTotalCommission(ctx)

	if err != nil {
		return nil, err
	}

	reply := pb.ListTotalCommissionReply{}
	for _, comm := range comms {
		reply.Commissions = append(reply.Commissions, &pb.GetUserTotalCommissionReply{
			Id:                     comm.Id,
			UserId:                 comm.UserId,
			TotalCommission:        comm.TotalCommission,
			TodayCommission:        comm.TodayCommission,
			SettledCommission:      comm.SettledCommission,
			TodayRegistrationCount: int32(comm.TodayRegistrationCount),
			TotalRegistrationCount: int32(comm.TotalRegistrationCount),
		})
	}
	return &reply, nil
}

func (s *CommissionService) ListTotalCommissionByParent(ctx context.Context, req *pb.ListTotalCommissionByParentReq) (*pb.ListTotalCommissionByParentReply, error) {
	comms, err := s.comm.ListTotalCommissionByParent(ctx, req.ParentId)

	if err != nil {
		return nil, err
	}

	reply := pb.ListTotalCommissionByParentReply{}
	for _, comm := range comms {
		reply.Commissions = append(reply.Commissions, &pb.GetUserTotalCommissionReply{
			Id:                     comm.Id,
			UserId:                 comm.UserId,
			TotalCommission:        comm.TotalCommission,
			TodayCommission:        comm.TodayCommission,
			SettledCommission:      comm.SettledCommission,
			TodayRegistrationCount: int32(comm.TodayRegistrationCount),
			TotalRegistrationCount: int32(comm.TotalRegistrationCount),
		})
	}
	return &reply, nil
}

func (s *CommissionService) ListCommissionByUser(ctx context.Context, req *pb.ListCommissionByUserReq) (*pb.ListCommissionByUserReply, error) {
	commissions, err := s.comm.ListCommissionByUser(ctx, req)
	if err != nil {
		return nil, err
	}

	reply := pb.ListCommissionByUserReply{UserId: req.UserId}
	for _, commission := range commissions {
		reply.Commissions = append(reply.Commissions, &pb.ListCommissionByUserReply_Commission{
			IndirectRechargeAmount:    commission.IndirectRechargeAmount,
			IndirectRegistrationCount: commission.IndirectRegistrationCount,
			DirectRechargeAmount:      commission.DirectRechargeAmount,
			DirectRegistrationCount:   commission.DirectRegistrationCount,
			Date:                      timestamppb.New(commission.Date),
		})
	}
	return &reply, nil
}
