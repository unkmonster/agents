package service

import (
	"context"

	pb "agents/api/authn/service/v1"
	"agents/app/authn/service/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
)

type AuthnService struct {
	pb.UnimplementedAuthnServer
	uc *biz.AuthUserCase
}

func NewAuthnService(uc *biz.AuthUserCase) *AuthnService {
	return &AuthnService{
		uc: uc,
	}
}

func (s *AuthnService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthReply, error) {
	if req.Password == nil {
		return nil, errors.New(400, "MISSING_PASSWORD", "缺少用户密码")
	}
	if req.Username == nil {
		return nil, errors.New(400, "MISSING_USER_NAME", "缺少用户名")
	}
	return s.uc.Login(ctx, req)
}

func (s *AuthnService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthReply, error) {
	return s.uc.RegisterChildUser(ctx, req)
}

func (s *AuthnService) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyReply, error) {
	panic("unimplement")
}
