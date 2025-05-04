package biz

import (
	pb "agents/api/user/service/v1"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type UserDomain struct {
	Id        string    `db:"id"`
	UserId    string    `db:"user_id"`
	Domain    string    `db:"domain"`
	CreatedAT time.Time `db:"created_at"`
}

type UserDomainRepo interface {
	Create(ctx context.Context, u *UserDomain) error
	Get(ctx context.Context, id string) (*UserDomain, error)
	List(ctx context.Context) ([]*UserDomain, error)
	ListByUserId(ctx context.Context, userId string) ([]*UserDomain, error)
	Delete(ctx context.Context, id string) error
}

type UserDomainUseCase struct {
	repo UserDomainRepo
	log  *log.Helper
}

func NewUserDomainUseCase(repo UserDomainRepo, logger log.Logger) *UserDomainUseCase {
	return &UserDomainUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserDomainUseCase) CreateUserDomain(ctx context.Context, req *pb.CreateUserDomainRequest) (*UserDomain, error) {
	if req.UserId == nil {
		return nil, errors.New(400, "MISSING_USER_ID", "缺少用户 ID")
	}

	if req.Domain == nil {
		return nil, errors.New(400, "MISSING_DOMAIN", "缺少域名")
	}

	userDomain := &UserDomain{
		Id:     uuid.New().String(),
		UserId: *req.UserId,
		Domain: *req.Domain,
	}

	if err := uc.repo.Create(ctx, userDomain); err != nil {
		return nil, err
	}

	return userDomain, nil
}

func (uc *UserDomainUseCase) Get(ctx context.Context, id string) (*UserDomain, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *UserDomainUseCase) ListAll(ctx context.Context) ([]*UserDomain, error) {
	return uc.repo.List(ctx)
}

func (uc *UserDomainUseCase) ListByUserId(ctx context.Context, userId string) ([]*UserDomain, error) {
	return uc.repo.ListByUserId(ctx, userId)
}

func (uc *UserDomainUseCase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
