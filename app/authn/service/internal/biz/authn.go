package biz

import (
	"context"
	"time"

	pb "agents/api/authn/service/v1"
	"agents/app/authn/service/internal/conf"
	"agents/pkg/encrypt"
	"agents/pkg/jwt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	mjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const MaxAgentLevel = 2

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
	comm CommissionRepo

	privKey       interface{}
	pubKey        interface{}
	tokenDuration time.Duration
	signMethod    string
}

func NewAuthUserCase(repo UserCredentialRepo, logger log.Logger, ur UserRepo, auth *conf.Auth, comm CommissionRepo) *AuthUserCase {
	var privKey interface{}
	var pubKey interface{}
	var err error

	if auth.SigningMethod == "RS256" {
		privKey, err = encrypt.LoadRsaPrivateKey(auth.Secret)
		if err == nil {
			pubKey, err = encrypt.LoadRSAPublicKey(auth.PublicKey)
		}
	} else {
		privKey = auth.Secret
		pubKey = auth.PublicKey
	}

	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	return &AuthUserCase{
		cr:   repo,
		log:  log.NewHelper(logger),
		ur:   ur,
		comm: comm,

		privKey:       privKey,
		pubKey:        pubKey,
		tokenDuration: auth.TokenDuration.AsDuration(),
		signMethod:    auth.SigningMethod,
	}
}

func (uc *AuthUserCase) generateTokenReply(user *User) (*pb.AuthReply, error) {
	// generate token
	tokenStr, err := jwt.GenerateJwt(&jwt.User{
		UserId: user.Id,
		Level:  int(user.Level),
	}, uc.privKey, time.Now().Add(uc.tokenDuration), uc.signMethod)

	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		Token: &tokenStr,
		User: &pb.UserInfo{
			Id:           &user.Id,
			Username:     &user.Username,
			Nickname:     user.Nickname,
			ParentId:     user.ParentId,
			Level:        &user.Level,
			SharePercent: user.SharePercent,
		},
	}, nil
}

// Register 不是注册，用于上级代理创建他的下级
func (uc *AuthUserCase) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthReply, error) {
	t, _ := mjwt.FromContext(ctx)
	token, ok := t.(*jwt.UserClaims)
	if !ok {
		return nil, errors.New(401, "INVALID_TOKEN", "无效 token")
	}

	if token.Subject != *req.ParentId {
		return nil, errors.New(403, "ILLEGAL_PARENT_ID", "非法父代理 ID")
	}

	// 被注册用户等级不允许小于调用者，并且不大于 MAX_LEVEL
	if *req.Level > MaxAgentLevel || *req.Level <= int32(token.Level) {
		return nil, errors.New(403, "ILLEGAL_USER_LEVEL", "不合法的代理等级")
	}

	user, err := uc.ur.Create(ctx, &User{
		Username:     *req.Username,
		Nickname:     req.Nickname,
		ParentId:     req.ParentId,
		Level:        *req.Level,
		SharePercent: req.SharePercent,
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
		HashedPassword: string(h),
	})
	if err != nil {
		return nil, err
	}

	err = uc.comm.InitUserCommission(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		User: &pb.UserInfo{
			Id:           &user.Id,
			Username:     &user.Username,
			Nickname:     user.Nickname,
			ParentId:     user.ParentId,
			Level:        &user.Level,
			SharePercent: user.SharePercent,
		},
	}, nil
}

func (uc *AuthUserCase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthReply, error) {
	credential, err := uc.cr.GetByUsername(ctx, *req.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(credential.HashedPassword), []byte(*req.Password)); err != nil {
		return nil, err
	}

	user, err := uc.ur.GetByUsername(ctx, *req.Username)
	if err != nil {
		return nil, err
	}
	return uc.generateTokenReply(user)
}
