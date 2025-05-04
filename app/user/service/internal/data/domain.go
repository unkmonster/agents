package data

import (
	"agents/app/user/service/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type userDomainRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserDomainRepo(data *Data, logger log.Logger) *userDomainRepo {
	return &userDomainRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *userDomainRepo) Create(ctx context.Context, domain *biz.UserDomain) error {
	query := `
		INSERT INTO user_domains (
			id,
			user_id,
			domain
		) VALUES (
			:id,
			:user_id,
			:domain 
		);
	`
	_, err := repo.data.db.NamedExecContext(ctx, query, domain)
	return err
}

func (repo *userDomainRepo) Get(ctx context.Context, id string) (*biz.UserDomain, error) {
	query := `
		SELECT *
		FROM user_domains
		WHERE id = ?;
	`
	dst := biz.UserDomain{}
	err := repo.data.db.GetContext(ctx, &dst, query, id)
	return &dst, err
}

func (repo *userDomainRepo) List(ctx context.Context) ([]*biz.UserDomain, error) {
	query := `
		SELECT *
		FROM user_domains;
	`
	dst := []*biz.UserDomain{}
	err := repo.data.db.SelectContext(ctx, &dst, query)
	return dst, err
}

func (repo *userDomainRepo) ListByUserId(ctx context.Context, userId string) ([]*biz.UserDomain, error) {
	query := `
		SELECT *
		FROM user_domains
		WHERE user_id = ?
	`
	dst := []*biz.UserDomain{}
	err := repo.data.db.SelectContext(ctx, &dst, query, userId)
	return dst, err
}

func (repo *userDomainRepo) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM user_domains
		WHERE id = ?;
	`
	_, err := repo.data.db.ExecContext(ctx, query, id)
	return err
}
