package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Wallet struct {
	Id         string    `db:"id"`
	UserId     string    `db:"user_id"`
	WalletType string    `db:"wallet_type"`
	Account    *string   `db:"account"`
	QrCode     *string   `db:"qr_code"`
	CreatedAt  time.Time `db:"created_at"`
}

type WalletRepo interface {
	Create(ctx context.Context, wallet *Wallet) (*Wallet, error)
	Update(ctx context.Context, wallet *Wallet) (*Wallet, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*Wallet, error)
	List(ctx context.Context) ([]*Wallet, error)
	ListByUser(ctx context.Context, userId string) ([]*Wallet, error)
}

type WalletUseCase struct {
	repo WalletRepo
	log  *log.Helper
}

func NewWalletUseCase(repo WalletRepo, logger log.Logger) *WalletUseCase {
	return &WalletUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *WalletUseCase) Create(ctx context.Context, wallet *Wallet) (*Wallet, error) {
	return uc.repo.Create(ctx, wallet)
}

func (uc *WalletUseCase) Update(ctx context.Context, wallet *Wallet) (*Wallet, error) {
	return uc.repo.Update(ctx, wallet)
}

func (uc *WalletUseCase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *WalletUseCase) Get(ctx context.Context, id string) (*Wallet, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *WalletUseCase) List(ctx context.Context) ([]*Wallet, error) {
	return uc.repo.List(ctx)
}

func (uc *WalletUseCase) ListByUser(ctx context.Context, userId string) ([]*Wallet, error) {
	return uc.repo.ListByUser(ctx, userId)
}
