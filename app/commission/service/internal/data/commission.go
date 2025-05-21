package data

import (
	"agents/app/commission/service/internal/biz"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	regTypeDirect   = "direct"
	regTypeIndirect = "indirect"
)

var _ biz.CommissionRepo = (*commissionRepo)(nil)

type commissionRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommissionRepo(data *Data, logger log.Logger) biz.CommissionRepo {
	return &commissionRepo{
		data: data,
		log:  log.NewHelper(logger),
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

// incUserTotalCommission 增加用户总佣金（累计佣金）
func (c *commissionRepo) incUserTotalCommission(ctx context.Context, tx *sqlx.Tx, userId string, amount int64) error {
	query := `
		INSERT INTO user_commissions (
			id,
			user_id,
			total_commission
		) VALUES (
		 	?,
			?,
			? 
		) ON DUPLICATE KEY
		 	UPDATE total_commission = total_commission + ?;
	`
	_, err := tx.ExecContext(ctx, query, uuid.NewString(), userId, amount, amount)
	return err
}

func (c *commissionRepo) incUserTotalRegistrationCount(ctx context.Context, tx *sqlx.Tx, userId string) error {
	query := `
		INSERT INTO user_commissions (
			id,
			user_id,
			total_registration_count
		) VALUES (
			?,
			?,
			?
		) ON DUPLICATE KEY
		 	UPDATE total_registration_count = total_registration_count + 1;
	`
	_, err := tx.ExecContext(ctx, query, uuid.NewString(), userId, 1)
	return err
}

func (c *commissionRepo) incUserCommission(ctx context.Context, userId string, amount int64, commType string) (err error) {
	var tx *sqlx.Tx
	tx, err = c.data.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				c.log.Errorf("rollback failed: %v", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				c.log.Errorf("commit failed: %v", err)
			}
		}
	}()

	if commType != biz.CommissionTypeDirect && commType != biz.CommissionTypeIndirect {
		return fmt.Errorf("invalid commission type: %s", commType)
	}

	query := `
		INSERT INTO daily_user_commissions (
			date,	
			user_id,
			%s_recharge_amount
		) VALUES (
			CURRENT_DATE,
			?,
			? 
		) ON DUPLICATE KEY
		 	UPDATE %s_recharge_amount = %s_recharge_amount + ?;
	`
	query = fmt.Sprintf(query, commType, commType, commType)

	_, err = tx.ExecContext(ctx, query, userId, amount, amount)
	if err != nil {
		return err
	}

	err = c.incUserTotalCommission(ctx, tx, userId, amount)
	return err
}

func (c *commissionRepo) IncUserDirectCommission(ctx context.Context, userId string, amount int64) (err error) {
	return c.incUserCommission(ctx, userId, amount, biz.CommissionTypeDirect)
}

func (c *commissionRepo) IncUserIndirectCommission(ctx context.Context, userId string, amount int64) error {
	return c.incUserCommission(ctx, userId, amount, biz.CommissionTypeIndirect)
}

func (c *commissionRepo) incUserRegistrationCount(ctx context.Context, userId string, regType string) (err error) {
	if regType != regTypeDirect && regType != regTypeIndirect {
		return fmt.Errorf("invalid registration type: %s", regType)
	}

	var tx *sqlx.Tx
	tx, err = c.data.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				c.log.Errorf("rollback failed: %v", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				c.log.Errorf("commit failed: %v", err)
			}
		}
	}()

	query := `
		INSERT INTO daily_user_commissions (
			date,
			user_id,
			%s_registration_count
		) VALUES (
			CURRENT_DATE,
			?,
			1
		) ON DUPLICATE KEY
		 	UPDATE %s_registration_count = %s_registration_count + 1;
	`
	query = fmt.Sprintf(query, regType, regType, regType)
	_, err = tx.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return c.incUserTotalRegistrationCount(ctx, tx, userId)
}

func (c *commissionRepo) IncUserDirectRegistrationCount(ctx context.Context, userId string) error {
	return c.incUserRegistrationCount(ctx, userId, regTypeDirect)
}

func (c *commissionRepo) IncUserIndirectRegistrationCount(ctx context.Context, userId string) error {
	return c.incUserRegistrationCount(ctx, userId, regTypeIndirect)
}
