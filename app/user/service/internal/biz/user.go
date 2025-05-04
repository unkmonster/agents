package biz

import (
	"context"
	"encoding/hex"
	"time"

	pb "agents/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             string    `db:"id"`
	Username       string    `db:"username"`
	hashedPassword string    `db:"hashed_password"`
	Nickname       *string   `db:"nickname"`
	ParentId       *string   `db:"parent_id"`
	Level          int32     `db:"level"`
	CreatedAt      time.Time `db:"created_at"`
}

// TODO: update, listByParentId
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
	// 仅允许 0 级代理（管理员）没有父级代理
	if *req.Level != 0 && req.ParentId == nil {
		return nil, errors.New(400, "MISSING_PARENT_ID", "缺少父级代理 ID ")
	}

	// TODO: parent_id 仅允许为调用者 ID, level 必须大于调用者 level 并且小于等于 max_level

	h, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Id:             uuid.New().String(),
		Username:       *req.Username,
		hashedPassword: hex.EncodeToString(h),
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
