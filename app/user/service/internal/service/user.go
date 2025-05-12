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
	if req.Username == nil {
		return nil, errors.New(400, "MISSING_USERNAME", "缺少用户名")
	}

	if req.Level == nil {
		return nil, errors.New(400, "MISSING_AGENT_LEVEL", "缺少代理等级")
	}

	return s.user.CreateUser(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	if req.Id == nil {
		return nil, errors.New(400, "MISSING_USER_ID", "缺少用户 ID")
	}

	return &pb.DeleteUserReply{}, s.user.DeleteUser(ctx, *req.Id)
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	if req.Id == nil {
		return nil, errors.New(400, "MISSING_USER_ID", "缺少用户 ID")
	}

	user, err := s.user.GetUser(ctx, *req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{
		Id:           &user.Id,
		Username:     &user.Username,
		Nickname:     user.Nickname,
		ParentId:     user.ParentId,
		Level:        &user.Level,
		SharePercent: &user.SharePercent,
	}, nil
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

func (s *UserService) GetUserDomain(ctx context.Context, req *pb.GetUserDomainRequest) (*pb.GetUserDomainReply, error) {
	if req.Id == nil {
		return nil, errors.New(400, "MISSING_DOMAIN_ID", "缺少域名 ID")
	}

	domain, err := s.domain.Get(ctx, *req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserDomainReply{
		Id:     &domain.Id,
		UserId: &domain.UserId,
		Domain: &domain.Domain,
	}, nil
}

func (s *UserService) ListUserDomains(ctx context.Context, req *pb.ListUserDomainsRequest) (*pb.ListUserDomainsReply, error) {
	domains, err := s.domain.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	reply := pb.ListUserDomainsReply{}
	for _, domain := range domains {
		reply.Domains = append(reply.Domains, &pb.ListUserDomainsReply_Domain{
			Id:     &domain.Id,
			UserId: &domain.UserId,
			Domain: &domain.Domain,
		})
	}
	return &reply, nil
}

func (s *UserService) ListUserDomainsByUserId(ctx context.Context, req *pb.ListUserDomainsByUserIdRequest) (*pb.ListUserDomainsByUserIdReply, error) {
	if req.UserId == nil {
		return nil, errors.New(400, "MISSING_USER_ID", "缺少用户 ID")
	}

	domains, err := s.domain.ListByUserId(ctx, *req.UserId)
	if err != nil {
		return nil, err
	}
	reply := pb.ListUserDomainsByUserIdReply{}
	for _, domain := range domains {
		reply.Domains = append(reply.Domains, &pb.ListUserDomainsByUserIdReply_Domain{
			Id:     &domain.Id,
			UserId: &domain.UserId,
			Domain: &domain.Domain,
		})
	}
	return &reply, nil

}

func (s *UserService) DeleteDomain(ctx context.Context, req *pb.DeleteDomainRequest) (*pb.DeleteDomainReply, error) {
	if req.Id == nil {
		return nil, errors.New(400, "MISSING_DOMAIN_ID", "缺少域名 ID")
	}

	return &pb.DeleteDomainReply{}, s.domain.Delete(ctx, *req.Id)
}

func (s *UserService) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserReply, error) {
	if req.Username == nil {
		return nil, errors.New(400, "MISSING_USER_NAME", "缺少用户名")
	}

	user, err := s.user.GetUserByUsername(ctx, *req.Username)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{
		Id:           &user.Id,
		Username:     &user.Username,
		Nickname:     user.Nickname,
		ParentId:     user.ParentId,
		Level:        &user.Level,
		SharePercent: &user.SharePercent,
	}, nil
}

func (s *UserService) GetUserByDomain(ctx context.Context, req *pb.GetUserByDomainRequest) (*pb.GetUserByDomainReply, error) {
	user, err := s.user.GetUserByDomain(ctx, req.Domain)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserByDomainReply{
		Id:           user.Id,
		Username:     user.Username,
		Nickname:     *user.Nickname,
		ParentId:     *user.ParentId,
		Level:        user.Level,
		SharePercent: user.SharePercent,
	}, nil
}
