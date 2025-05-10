package biz

import (
	"context"
	"time"
)

type User struct {
	Id           string    `db:"id"`
	Username     string    `db:"username"`
	Nickname     *string   `db:"nickname"`
	ParentId     *string   `db:"parent_id"`
	Level        int32     `db:"level"`
	CreatedAt    time.Time `db:"created_at"`
	SharePercent float32   `db:"share_percent"`
}

type UserRepo interface {
	GetUser(ctx context.Context, userId string) (*User, error)
	GetUserByDomain(ctx context.Context, domain string) (*User, error)
}
