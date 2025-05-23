package biz

import (
	"context"
	"database/sql"
	"time"

	pb "agents/api/user/service/v1"
	"agents/pkg/mysql"
	"agents/pkg/paging"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	Id           string       `db:"id"`
	Username     string       `db:"username"`
	Nickname     *string      `db:"nickname"`
	ParentId     *string      `db:"parent_id"`
	Level        int32        `db:"level"`
	CreatedAt    time.Time    `db:"created_at"`
	SharePercent float32      `db:"share_percent"`
	LastLoginAt  sql.NullTime `db:"last_login_at"`
}

// TODO: update, listByParentId
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByDomain(ctx context.Context, domain string) (*User, error)
	ListUserByParent(ctx context.Context, parentId string, paging *paging.Paging) ([]*User, error)
	GetZeroUser(ctx context.Context) (*User, error)
	UpdateUserLastLoginTime(ctx context.Context, id string) error
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
	if req.Level != 0 && req.ParentId == nil {
		return nil, pb.ErrorMissingParentId("")
	}

	user := User{
		Id:           uuid.New().String(),
		Username:     req.Username,
		Nickname:     req.Nickname,
		ParentId:     req.ParentId,
		Level:        req.Level,
		SharePercent: req.SharePercent,
	}

	if req.Level == 0 {
		user.Id = "0000000-0000-0000-0000-000000000000"
	}

	if err := uc.repo.CreateUser(ctx, &user); err != nil {
		uc.log.Infof("%#v", err)
		if mysql.IsDuplicateEntryError(err) {
			return nil, pb.ErrorUserIsExists("")
		}
		return nil, err
	}

	return &pb.CreateUserReply{
		User: &pb.UserInfo2{
			Id:           user.Id,
			Username:     user.Username,
			Nickname:     user.Nickname,
			ParentId:     user.ParentId,
			Level:        (user.Level),
			SharePercent: user.SharePercent,
			CreatedAt:    timestamppb.Now(),
		},
	}, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	return uc.repo.GetUserByUsername(ctx, username)
}

func (uc *UserUseCase) GetUserByDomain(ctx context.Context, domain string) (*User, error) {
	return uc.repo.GetUserByDomain(ctx, domain)
}

func (uc *UserUseCase) ListUserByParent(ctx context.Context, parentId string, paging *paging.Paging) ([]*User, error) {
	return uc.repo.ListUserByParent(ctx, parentId, paging)
}

func (uc *UserUseCase) UpdateUserLastLoginTime(ctx context.Context, id string) error {
	return uc.repo.UpdateUserLastLoginTime(ctx, id)
}
