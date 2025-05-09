package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	Id          string    `db:"id"`
	PaymentType string    `db:"payment_type"`
	Name        string    `db:"name"`
	Amount      int32     `db:"amount"`
	Domain      string    `db:"domain"`
	CreatedAt   time.Time `db:"created_at"`
}

type OrderRepo interface {
	Create(ctx context.Context, order *Order) (*Order, error)
	Get(ctx context.Context, id string) (*Order, error)
	List(ctx context.Context) ([]*Order, error)
	ListByUser(ctx context.Context, userId string) ([]*Order, error)
	ListByDomain(ctx context.Context, domain string) ([]*Order, error)
}

type OrderUseCase struct {
	log  *log.Helper
	repo OrderRepo
}

func NewOrderUseCase(logger log.Logger, repo OrderRepo) *OrderUseCase {
	return &OrderUseCase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (uc *OrderUseCase) Create(ctx context.Context, order *Order) (*Order, error) {
	return uc.repo.Create(ctx, order)
}

func (uc *OrderUseCase) Get(ctx context.Context, id string) (*Order, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *OrderUseCase) List(ctx context.Context) ([]*Order, error) {
	return uc.repo.List(ctx)
}

func (uc *OrderUseCase) ListByUser(ctx context.Context, userId string) ([]*Order, error) {
	return uc.repo.ListByUser(ctx, userId)
}

func (uc *OrderUseCase) ListByDomain(ctx context.Context, domain string) ([]*Order, error) {
	return uc.repo.ListByDomain(ctx, domain)
}
