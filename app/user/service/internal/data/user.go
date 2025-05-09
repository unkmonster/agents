package data

import (
	"agents/app/user/service/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	query := `
		INSERT INTO users (
			id,
			username,
			nickname,
			parent_id,
			level
		) VALUES (
			:id,
			:username,
			:nickname,
			:parent_id,
			:level
		);
	`
	_, err := r.data.db.NamedExecContext(ctx, query, user)
	return err
}

func (r *userRepo) GetUser(ctx context.Context, id string) (*biz.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE id = ?;
	`
	dst := biz.User{}
	err := r.data.db.GetContext(ctx, &dst, query, id)
	return &dst, err
}

func (r *userRepo) DeleteUser(ctx context.Context, id string) error {
	query := `
		DELETE FROM users
		WHERE id = ?;
	`
	_, err := r.data.db.ExecContext(ctx, query, id)
	return err
}

// GetUserByUsername implements biz.UserRepo.
func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE username = ?;
	`
	dst := biz.User{}

	err := r.data.db.GetContext(ctx, &dst, query, username)
	return &dst, err
}

func (r *userRepo) GetUserByDomain(ctx context.Context, domain string) (*biz.User, error) {
	query := `
		SELECT users.*
		FROM users
		INNER JOIN user_domains
		ON users.id = user_domains.user_id
		WHERE user_domains.domain = ?;
	`
	dst := biz.User{}
	err := r.data.db.GetContext(ctx, &dst, query, domain)
	return &dst, err
}
