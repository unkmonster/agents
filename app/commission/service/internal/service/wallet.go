package service

import (
	"context"

	pb "agents/api/commission/service/v1"
	"agents/app/commission/service/internal/biz"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type WalletService struct {
	pb.UnimplementedWalletServer
	wallet *biz.WalletUseCase
}

func NewWalletService(wallet *biz.WalletUseCase) *WalletService {
	return &WalletService{
		wallet: wallet,
	}
}

func (s *WalletService) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletReply, error) {
	wallet, err := s.wallet.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateWalletReply{
		UserId:     wallet.UserId,
		Account:    wallet.Account,
		QrCode:     wallet.QrCode,
		Id:         wallet.Id,
		WalletType: wallet.WalletType,
		CreatedAt:  timestamppb.New(wallet.CreatedAt),
	}, nil
}

func (s *WalletService) UpdateWallet(ctx context.Context, req *pb.UpdateWalletRequest) (*pb.UpdateWalletReply, error) {
	wallet, err := s.wallet.Update(ctx, &biz.Wallet{
		Id:      req.Id,
		QrCode:  req.QrCode,
		Account: req.Account,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateWalletReply{
		UserId:     wallet.UserId,
		Account:    wallet.Account,
		QrCode:     wallet.QrCode,
		Id:         wallet.Id,
		WalletType: wallet.WalletType,
		CreatedAt:  timestamppb.New(wallet.CreatedAt),
	}, nil
}
func (s *WalletService) DeleteWallet(ctx context.Context, req *pb.DeleteWalletRequest) (*pb.DeleteWalletReply, error) {
	err := s.wallet.Delete(ctx, req.Id)
	return nil, err
}

func (s *WalletService) GetWallet(ctx context.Context, req *pb.GetWalletRequest) (*pb.GetWalletReply, error) {
	wallet, err := s.wallet.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetWalletReply{
		UserId:     wallet.UserId,
		Account:    wallet.Account,
		QrCode:     wallet.QrCode,
		Id:         wallet.Id,
		WalletType: wallet.WalletType,
		CreatedAt:  timestamppb.New(wallet.CreatedAt),
	}, nil
}

func (s *WalletService) ListWallet(ctx context.Context, req *pb.ListWalletRequest) (*pb.ListWalletReply, error) {
	wallets, err := s.wallet.List(ctx)
	if err != nil {
		return nil, err
	}

	reply := pb.ListWalletReply{}
	for _, wallet := range wallets {
		reply.Wallets = append(reply.Wallets, &pb.GetWalletReply{
			UserId:     wallet.UserId,
			Account:    wallet.Account,
			QrCode:     wallet.QrCode,
			Id:         wallet.Id,
			WalletType: wallet.WalletType,
			CreatedAt:  timestamppb.New(wallet.CreatedAt),
		})
	}
	return &reply, nil
}

func (s *WalletService) ListWalletByUser(ctx context.Context, req *pb.ListWalletByUserRequest) (*pb.ListWalletReply, error) {
	wallets, err := s.wallet.ListByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	reply := pb.ListWalletReply{}
	for _, wallet := range wallets {
		reply.Wallets = append(reply.Wallets, &pb.GetWalletReply{
			UserId:     wallet.UserId,
			Account:    wallet.Account,
			QrCode:     wallet.QrCode,
			Id:         wallet.Id,
			WalletType: wallet.WalletType,
			CreatedAt:  timestamppb.New(wallet.CreatedAt),
		})
	}
	return &reply, nil
}
