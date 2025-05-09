package data

import (
	"agents/app/order/service/internal/biz"
	"context"
)

var _ biz.OrderRepo = (*orderRepo)(nil)

type orderRepo struct {
	data *Data
}

// Create implements biz.OrderRepo.
func (o *orderRepo) Create(ctx context.Context, order *biz.Order) (*biz.Order, error) {
	query := `
		INSERT INTO orders (
			id,
			payment_type,
			name,
			amount,
			domain
		) VALUES (
			:id,
			:payment_type,
			:name,
			:amount,
			:domain
		);
	`
	_, err := o.data.db.NamedExecContext(ctx, query, order)
	return order, err
}

// Get implements biz.OrderRepo.
func (o *orderRepo) Get(ctx context.Context, id string) (*biz.Order, error) {
	query := `
		SELECT *
		FROM orders
		WHERE id = ?;
	`
	dst := biz.Order{}
	err := o.data.db.GetContext(ctx, &dst, query, id)
	return &dst, err
}

// List implements biz.OrderRepo.
func (o *orderRepo) List(ctx context.Context) ([]*biz.Order, error) {
	query := `
		SELECT *
		FROM orders;
	`
	dst := []*biz.Order{}
	err := o.data.db.SelectContext(ctx, &dst, query)
	return dst, err
}

// ListByDomain implements biz.OrderRepo.
func (o *orderRepo) ListByDomain(ctx context.Context, domain string) ([]*biz.Order, error) {
	query := `
		SELECT *
		FROM orders
		WHERE domain = ?;
	`
	dst := []*biz.Order{}
	err := o.data.db.SelectContext(ctx, &dst, query, domain)
	return dst, err
}

// ListByUser implements biz.OrderRepo.
func (o *orderRepo) ListByUser(ctx context.Context, userId string) ([]*biz.Order, error) {
	query := `
		SELECT *
		FROM orders
		WHERE user_id = ?;
	`
	dst := []*biz.Order{}
	err := o.data.db.SelectContext(ctx, &dst, query, userId)
	return dst, err
}

func NewOrderRepo(data *Data) biz.OrderRepo {
	return &orderRepo{
		data: data,
	}
}
