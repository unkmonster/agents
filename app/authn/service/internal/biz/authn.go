package biz

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	pb "agents/api/authn/service/v1"
	"agents/app/authn/service/internal/conf"
	"agents/pkg/encrypt"
	"agents/pkg/jwt"

	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	mjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const MaxAgentLevel = 2
const DefaultSigningAlg = "RS256"

type UserCredential struct {
	Id             string    `db:"id"`
	Username       string    `db:"username"`
	UserId         string    `db:"user_id"`
	HashedPassword string    `db:"hashed_password"`
	Alg            string    `db:"alg"`    // 总是 RS256
	Secret         *string   `db:"secret"` // 未使用
	PublicKey      *string   `db:"public_key"`
	PrivateKey     *string   `db:"private_key"`
	TokenKey       string    `db:"token_key"`
	CreatedAt      time.Time `db:"created_at"`
}

type UserCredentialRepo interface {
	Create(ctx context.Context, uc *UserCredential) (*UserCredential, error)
	GetByUsername(ctx context.Context, username string) (*UserCredential, error)
	GetByUserId(ctx context.Context, userId string) (*UserCredential, error)
}

type AuthUseCase struct {
	credential UserCredentialRepo
	user       UserRepo
	log        *log.Helper

	tokenDuration time.Duration

	gateway GatewayRepo
}

func NewAuthUserCase(repo UserCredentialRepo, logger log.Logger, ur UserRepo, auth *conf.Auth, gateway GatewayRepo) *AuthUseCase {
	return &AuthUseCase{
		credential: repo,
		log:        log.NewHelper(logger),
		user:       ur,

		tokenDuration: auth.TokenDuration.AsDuration(),
		gateway:       gateway,
	}
}

// 解析 PEM 格式的 RSA 私钥
func parseRSAPrivateKeyFromPEM(pemData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing RSA private key")
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

// doRegisterUser 做实际注册用户需要的工作，调用用户服务创建用户，保存用户凭据，创建网关消费者...
func (uc *AuthUseCase) doRegisterUser(ctx context.Context, password string, user *User) (*User, error) {
	username := user.Username

	// 调用用户服务创建用户
	uc.log.Infow("msg", "start creating user in user_service", "username", user.Username)
	user, err := uc.user.Create(ctx, user)
	if err != nil {
		uc.log.Errorw("msg", "creating user in user_service failed", "username", username, "reason", err)
		return nil, err
	}

	// 创建网关消费者
	uc.log.Infow("msg", "start creating consumer in gateway", "username", user.Username)
	consumer, err := uc.gateway.CreateUserConsumer(ctx, user)
	if err != nil {
		uc.log.Errorw("msg", "creating consumer in gateway failed", "username", user.Username, "reason", err)
		return nil, err
	}

	// 生成一对 RSA 密钥
	priv, pub, err := encrypt.GenerateRS256KeyPair()
	if err != nil {
		uc.log.Errorw("msg", "generation RSA key-pair failed", "username", user.Username, "reason", err)
		return nil, err
	}

	// 创建凭据为网关消费者
	uc.log.Infow("msg", "start creating credential for gateway consumer", "consumer", consumer)
	jwtKey, err := uc.gateway.CreateConsumerCredential(ctx, consumer, pub)
	if err != nil {
		uc.log.Errorw("msg", "creating credential for gateway consumer failed", "username", user.Username, "reason", err)
		return nil, err
	}

	// 保存用户凭据
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		uc.log.Errorw("hash user password failed", "username", user.Username, "reason", err)
		return nil, err
	}

	uc.log.Infow("msg", "start saving user credentials", "username", user.Username, "userid", user.Id)
	_, err = uc.credential.Create(ctx, &UserCredential{
		Id:             uuid.NewString(),
		Username:       user.Username,
		UserId:         user.Id,
		HashedPassword: string(h),
		Alg:            "RS256",
		PublicKey:      &pub,
		PrivateKey:     &priv,
		TokenKey:       jwtKey,
		CreatedAt:      time.Now(),
	})
	if err != nil {
		uc.log.Errorw("msg", "saving user credential failed", "username", user.Username, "reason", err)
		return nil, err
	}
	return user, nil
}

func (uc *AuthUseCase) RegisterZeroUser(ctx context.Context, username, password string) (*User, error) {
	return uc.doRegisterUser(ctx, password, &User{
		Username:     username,
		Level:        0,
		SharePercent: 1,
		ParentId:     nil,
	})
}

// RegisterChildUser 不是注册，用于上级代理创建他的下级
func (uc *AuthUseCase) RegisterChildUser(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthReply, error) {
	t, _ := mjwt.FromContext(ctx)
	token, ok := t.(*jwt.UserClaims)
	if !ok {
		return nil, kerrors.New(401, "INVALID_TOKEN", "无效 token")
	}

	if token.Subject != req.ParentId {
		return nil, pb.ErrorIllegalParentId("非法操作")
	}

	// 被注册用户等级不允许小于调用者，并且不大于 MAX_LEVEL
	if req.Level > MaxAgentLevel || req.Level <= int32(token.Level) {
		return nil, pb.ErrorIllegalUserLevel("用户等级不合法: %d", req.Level)
	}

	user, err := uc.doRegisterUser(ctx, req.Password, &User{
		Username:     req.Username,
		Nickname:     req.Nickname,
		ParentId:     &req.ParentId,
		Level:        req.Level,
		SharePercent: req.SharePercent,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		User: &pb.UserInfo{
			Id:           user.Id,
			Username:     user.Username,
			Nickname:     user.Nickname,
			ParentId:     user.ParentId,
			Level:        user.Level,
			SharePercent: user.SharePercent,
			CreatedAt:    timestamppb.New(user.CreatedAt),
		},
	}, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthReply, error) {
	credential, err := uc.credential.GetByUsername(ctx, *req.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(credential.HashedPassword), []byte(*req.Password)); err != nil {
		return nil, err
	}

	user, err := uc.user.GetByUsername(ctx, *req.Username)
	if errors.Is(err, sql.ErrNoRows) {
		err = pb.ErrorUserNotFount("user %q not fount", *req.Username)
	}
	if err != nil {
		return nil, err
	}

	// pem -> *rsa.PrivateKey
	pk, err := parseRSAPrivateKeyFromPEM([]byte(*credential.PrivateKey))
	if err != nil {
		return nil, err
	}

	tokenStr, err := jwt.GenerateJwt(jwt.UserClaims{
		RegisteredClaims: jwtv5.RegisteredClaims{
			Subject:   user.Id,
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(uc.tokenDuration)),
			NotBefore: jwtv5.NewNumericDate(time.Now()),
			Issuer:    credential.TokenKey,
		},
		Level: int(user.Level),
	}, pk)
	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		Token: &tokenStr,
	}, nil
}

func (uc *AuthUseCase) GetUserCredential(ctx context.Context, userId string) (*UserCredential, error) {
	return uc.credential.GetByUserId(ctx, userId)
}

//func (uc *AuthUserCase) GetZeroUser
