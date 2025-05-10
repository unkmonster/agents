package biz

import "context"

type CommissionRepo interface {
	InitUserCommission(ctx context.Context, userId string) error
}
