package biz

import (
	"context"
	"encoding/hex"
	"time"

	pb "agents/api/authn/service/v1"
	"agents/app/authn/service/internal/conf"
	"agents/pkg/jwt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserCredential struct {
	Id             string    `db:"id"`
	Username       string    `db:"username"`
	UserId         string    `db:"user_id"`
	HashedPassword string    `db:"hashed_password"`
	CreatedAt      time.Time `db:"created_at"`
}

type UserCredentialRepo interface {
	Create(ctx context.Context, uc *UserCredential) (*UserCredential, error)
	GetByUsername(ctx context.Context, username string) (*UserCredential, error)
}

type AuthUserCase struct {
	cr   UserCredentialRepo
	ur   UserRepo
	log  *log.Helper
	auth *conf.Auth
	comm CommissionRepo
}

func NewAuthUserCase(repo UserCredentialRepo, logger log.Logger, ur UserRepo, auth *conf.Auth, comm CommissionRepo) *AuthUserCase {
	return &AuthUserCase{
		cr:   repo,
		log:  log.NewHelper(logger),
		ur:   ur,
		auth: auth,
		comm: comm,
	}
}

func (uc *AuthUserCase) generateTokenReply(user *User) (*pb.AuthReply, error) {
	// generate token
	tokenStr, err := jwt.GenerateJwt(&jwt.User{
		UserId: user.Id,
		Level:  int(user.Level),
	}, *uc.auth.JwtSecret, time.Now().Add(uc.auth.TokenTtl.AsDuration()))

	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		Token: &tokenStr,
		User: &pb.UserInfo{
			Id:       &user.Id,
			Username: &user.Username,
			Nickname: user.Nickname,
			ParentId: user.ParentId,
			Level:    &user.Level,
		},
	}, nil
}

func (uc *AuthUserCase) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthReply, error) {
	// TODO 被注册用户等级不允许小于调用者，并且不大于 MAX_LEVEL
	user, err := uc.ur.Create(ctx, &User{
		Username: *req.Username,
		Nickname: req.Nickname,
		ParentId: req.ParentId,
		Level:    *req.Level,
	})
	if err != nil {
		return nil, err
	}

	// hash password
	h, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = uc.cr.Create(ctx, &UserCredential{
		Id:             uuid.NewString(),
		UserId:         user.Id,
		Username:       user.Username,
		HashedPassword: hex.EncodeToString(h),
	})
	if err != nil {
		return nil, err
	}

	err = uc.comm.InitUserCommission(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	// generate token
	return uc.generateTokenReply(user)
}

func (uc *AuthUserCase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthReply, error) {
	credential, err := uc.cr.GetByUsername(ctx, *req.Username)
	if err != nil {
		return nil, err
	}

	h, err := hex.DecodeString(credential.HashedPassword)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(h, []byte(*req.Password)); err != nil {
		return nil, err
	}

	user, err := uc.ur.GetByUsername(ctx, *req.Username)
	if err != nil {
		return nil, err
	}
	return uc.generateTokenReply(user)
}

func (uc *AuthUserCase) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyReply, error) {
	_, err := jwt.ParseJWT(*req.Token, *uc.auth.JwtSecret)
	return &pb.VerifyReply{}, err // TODO 暂不返回用户信息
}
