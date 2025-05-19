package data

import (
	"agents/app/authn/service/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserCredentialRepo = (*userCredentialRepo)(nil)

type userCredentialRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserCredentialRepo(data *Data, logger log.Logger) biz.UserCredentialRepo {
	return &userCredentialRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create implements biz.UserCredentialRepo.
func (u *userCredentialRepo) Create(ctx context.Context, uc *biz.UserCredential) (*biz.UserCredential, error) {
	query := `
		INSERT INTO user_credentials (
			id,
			username,
			user_id,
			hashed_password,
			alg,
			public_key,
			private_key,
			token_key
		) VALUES (
			:id,
			:username,
			:user_id,
			:hashed_password,
			:alg,
			:public_key,
			:private_key,
			:token_key
		)
	`
	_, err := u.data.db.NamedExecContext(ctx, query, uc)
	return uc, err
}

// GetByUsername implements biz.UserCredentialRepo.
func (u *userCredentialRepo) GetByUsername(ctx context.Context, username string) (*biz.UserCredential, error) {
	query := `
		SELECT *
		FROM user_credentials
		WHERE username = ?;
	`
	dst := biz.UserCredential{}
	err := u.data.db.GetContext(ctx, &dst, query, username)
	return &dst, err
}
