package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"

	pb "agents/api/user/service/v1"
	"agents/app/user/service/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	user   *biz.UserUseCase
	domain *biz.UserDomainUseCase
}

func NewUserService(user *biz.UserUseCase, domain *biz.UserDomainUseCase) *UserService {
	return &UserService{
		user:   user,
		domain: domain,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func (s *UserService) CreateUserDomain(ctx context.Context, req *pb.CreateUserDomainRequest) (*pb.CreateUserDomainReply, error) {
	if req.UserId == nil {
		return nil, errors.New(400, "MISSING_USER_ID", "缺少用户 ID")
	}

	if req.Domain == nil {
		return nil, errors.New(400, "MISSING_DOMAIN", "缺少域名")
	}

	domain, err := s.domain.CreateUserDomain(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserDomainReply{
		Id:     &domain.Id,
		UserId: &domain.UserId,
		Domain: &domain.Domain,
	}, nil
}
