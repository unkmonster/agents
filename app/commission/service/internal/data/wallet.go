package data

import (
	"agents/app/commission/service/internal/biz"
	"context"
)

var _ biz.WalletRepo = (*walletRepo)(nil)

type walletRepo struct {
	data *Data
}

func NewWalletRepo(data *Data) biz.WalletRepo {
	return &walletRepo{
		data: data,
	}
}

// Create implements biz.WalletRepo.
func (w *walletRepo) Create(ctx context.Context, wallet *biz.Wallet) (*biz.Wallet, error) {
	query := `
		INSERT INTO user_wallets (
			id,
			user_id,
			wallet_type,
			account,
			qr_code
		) VALUES (
			:id,
			:user_id,
			:wallet_type,
			:account,
			:qr_code
		);
	`
	_, err := w.data.db.NamedExecContext(ctx, query, wallet)
	return wallet, err
}

// Delete implements biz.WalletRepo.
func (w *walletRepo) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM user_wallets
		WHERE id = ?;
	`
	_, err := w.data.db.ExecContext(ctx, query, id)
	return err
}

// Get implements biz.WalletRepo.
func (w *walletRepo) Get(ctx context.Context, id string) (*biz.Wallet, error) {
	query := `
		SELECT *
		FROM user_wallets
		WHERE id = ?;
	`
	dst := biz.Wallet{}
	err := w.data.db.GetContext(ctx, &dst, query, id)
	return &dst, err
}

// List implements biz.WalletRepo.
func (w *walletRepo) List(ctx context.Context) ([]*biz.Wallet, error) {
	query := `
		SELECT *
		FROM user_wallets;
	`
	dst := []*biz.Wallet{}
	err := w.data.db.SelectContext(ctx, &dst, query)
	return dst, err
}

// ListByUser implements biz.WalletRepo.
func (w *walletRepo) ListByUser(ctx context.Context, userId string) ([]*biz.Wallet, error) {
	query := `
		SELECT *
		FROM user_wallets
		WHERE user_id = ?;
	`
	dst := []*biz.Wallet{}
	err := w.data.db.SelectContext(ctx, &dst, query, userId)
	return dst, err
}

// Update implements biz.WalletRepo.
func (w *walletRepo) Update(ctx context.Context, wallet *biz.Wallet) (*biz.Wallet, error) {
	panic("unimplemented")
}
