package service

import (
	"context"

	pb "agents/api/authn/service/v1"
)

type AuthnService struct {
	pb.UnimplementedAuthnServer
}

func NewAuthnService() *AuthnService {
	return &AuthnService{}
}

func (s *AuthnService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthReply, error) {
	return &pb.AuthReply{}, nil
}
func (s *AuthnService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthReply, error) {
	return &pb.AuthReply{}, nil
}
func (s *AuthnService) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyReply, error) {
	return &pb.VerifyReply{}, nil
}
