package biz

import (
	"context"

	pb "agents/api/user/service/v1"

	errors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type User struct {
	Id             string
	Username       string
	HashedPassword string
	Nickname       *string
	ParentId       *string
	Level          int32
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	if req.Username == nil {
		return nil, errors.New(400, "MISSING_USERNAME", "缺少用户名")
	}

	if req.Password == nil {
		return nil, errors.New(400, "MISSING_PASSWORD", "缺少密码")
	}

	if req.Level == nil {
		return nil, errors.New(400, "MISSING_AGENT_LEVEL", "缺少代理等级")
	}

	// 仅允许 0 级代理（管理员）没有父级代理
	if *req.Level != 0 && req.ParentId == nil {
		return nil, errors.New(400, "MISSING_PARENT_ID", "缺少父级代理 ID ")
	}

	// TODO: parent_id 仅允许为调用者 ID, level 必须大于调用者 level 并且小于等于 max_level

	user := User{
		Id:             uuid.New().String(),
		Username:       *req.Username,
		HashedPassword: *req.Password,
		Nickname:       req.Nickname,
		ParentId:       req.ParentId,
		Level:          *req.Level,
	}

	if err := uc.repo.CreateUser(ctx, &user); err != nil {
		return nil, err
	}

	return &pb.CreateUserReply{
		Id:       &user.Id,
		Username: &user.Username,
		Nickname: user.Nickname,
		ParentId: user.ParentId,
		Level:    &(user.Level),
	}, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	return uc.repo.DeleteUser(ctx, id)
}
