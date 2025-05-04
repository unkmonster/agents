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
			hashed_password,
			nickname,
			parent_id,
			level
		) VALUES (
			:id,
			:username,
			:hashed_password,
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
