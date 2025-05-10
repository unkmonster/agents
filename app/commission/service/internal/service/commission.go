package service

import (
	"context"

	pb "agents/api/commission/service/v1"
	"agents/app/commission/service/internal/biz"
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

func (s *CommissionService) GetUserCommission(ctx context.Context, req *pb.GetUserCommissionRequest) (*pb.GetUserCommissionReply, error) {
	comm, err := s.comm.GetUserCommission(ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	return &pb.GetUserCommissionReply{
		Id:                comm.Id,
		UserId:            comm.UserId,
		TotalCommission:   comm.TotalCommission,
		TodayCommission:   comm.TodayCommission,
		SettledCommission: comm.SettledCommission,
	}, nil
}

func (s *CommissionService) ListCommission(ctx context.Context, req *pb.ListCommissionRequest) (*pb.ListCommissionReply, error) {
	comms, err := s.comm.ListCommission(ctx)

	if err != nil {
		return nil, err
	}

	reply := pb.ListCommissionReply{}
	for _, comm := range comms {
		reply.Commissions = append(reply.Commissions, &pb.GetUserCommissionReply{
			Id:                comm.Id,
			UserId:            comm.UserId,
			TotalCommission:   comm.TotalCommission,
			TodayCommission:   comm.TodayCommission,
			SettledCommission: comm.SettledCommission,
		})
	}
	return &reply, nil
}

func (s *CommissionService) ListCommissionByParent(ctx context.Context, req *pb.ListCommissionByParentReq) (*pb.ListCommissionByParentReply, error) {
	comms, err := s.comm.ListCommissionByParent(ctx, req.ParentId)

	if err != nil {
		return nil, err
	}

	reply := pb.ListCommissionByParentReply{}
	for _, comm := range comms {
		reply.Commissions = append(reply.Commissions, &pb.GetUserCommissionReply{
			Id:                comm.Id,
			UserId:            comm.UserId,
			TotalCommission:   comm.TotalCommission,
			TodayCommission:   comm.TodayCommission,
			SettledCommission: comm.SettledCommission,
		})
	}
	return &reply, nil
}
