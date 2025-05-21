package data

import (
	"agents/app/commission/service/internal/biz"
	"context"

	"github.com/google/uuid"
)

var _ biz.CommissionRepo = (*commissionRepo)(nil)

type commissionRepo struct {
	data *Data
}

func NewCommissionRepo(data *Data) biz.CommissionRepo {
	return &commissionRepo{
		data: data,
	}
}

// GetUserCommission implements biz.CommissionRepo.
func (c *commissionRepo) GetUserCommission(ctx context.Context, userId string) (*biz.Commission, error) {
	query := `
		SELECT *
		FROM user_commissions
		WHERE user_id = ?;
	`
	dst := biz.Commission{}
	err := c.data.db.GetContext(ctx, &dst, query, userId)
	return &dst, err
}

// IncUserCommission implements biz.CommissionRepo.
func (c *commissionRepo) IncUserCommission(ctx context.Context, userId string, amount int32) error {
	query := `
		UPDATE user_commissions
		SET today_commission = ? + today_commission,
			total_commission = ? + total_commission
		WHERE user_id = ?;
	`
	_, err := c.data.db.ExecContext(ctx, query, amount, amount, userId)
	return err
}

// IncUserSettledCommission implements biz.CommissionRepo.
func (c *commissionRepo) IncUserSettledCommission(ctx context.Context, userId string, amount int32) error {
	query := `
		UPDATE user_commissions
		SET settled_commission = ? + settled_commission
		WHERE user_id = ?;
	`
	_, err := c.data.db.ExecContext(ctx, query, amount, userId)
	return err
}

// InitUserCommission implements biz.CommissionRepo.
func (c *commissionRepo) InitUserCommission(ctx context.Context, userId string) error {
	query := `
		INSERT INTO user_commissions (
			id,
			user_id,
			today_commission,
			total_commission,
			settled_commission
		) VALUES (
			?,
			?,
			?,
			?,
			?
		);
	`
	_, err := c.data.db.ExecContext(ctx, query, uuid.NewString(), userId, 0, 0, 0)
	return err
}

// ListCommission implements biz.CommissionRepo.
func (c *commissionRepo) ListCommission(ctx context.Context) ([]*biz.Commission, error) {
	query := `
		SELECT *
		FROM user_commissions;
	`
	dst := []*biz.Commission{}
	err := c.data.db.SelectContext(ctx, &dst, query)
	return dst, err
}

// ListCommissionByParent implements biz.CommissionRepo.
func (c *commissionRepo) ListCommissionByParent(ctx context.Context, parentId string) ([]*biz.Commission, error) {
	panic("unimplemented")
}

func (c *commissionRepo) IncUserRegistrationCount(ctx context.Context, userId string) error {
	query := `
		UPDATE user_commissions
		SET today_registration_count = today_registration_count + 1,
			total_registration_count = total_registration_count + 1
		WHERE user_id = ?;
	`
	_, err := c.data.db.ExecContext(ctx, query, userId)
	return err
}
