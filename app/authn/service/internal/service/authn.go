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
	if req.Username == nil {
		return nil, errors.New(400, "MISSING_USER_NAME", "缺少用户名")
	}
	if req.Password == nil {
		return nil, errors.New(400, "MISSING_PASSWORD", "缺少用户密码")
	}
	if req.Level == nil {
		return nil, errors.New(400, "MISSING_LEVEL", "缺少等级")
	}
	if req.ParentId == nil {
		return nil, errors.New(400, "MISSING_PARENT_ID", "缺少 parent_id")
	}
	return s.uc.Register(ctx, req)
}

func (s *AuthnService) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyReply, error) {
	if req.Token == nil {
		return nil, errors.New(400, "MISSING_TOKEN", "缺少 token")
	}
	return s.uc.Verify(ctx, req)
}
