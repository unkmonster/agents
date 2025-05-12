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
	SharePercent float32
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
}
