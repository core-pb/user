package main

import (
	"context"

	"github.com/core-pb/dt/time/v1"
	"github.com/core-pb/user/user/v1"
	"github.com/redis/rueidis"
	"github.com/uptrace/bun"
)

var (
	db    *bun.DB
	cache rueidis.Client
)

type User struct {
	bun.BaseModel `bun:"table:user"`
	*user.User
}

type UserAuth struct {
	bun.BaseModel `bun:"table:user_auth"`
	*user.UserAuth
}

type UserTag struct {
	bun.BaseModel `bun:"table:user_tag"`
	*user.UserTag
}

type UserDetail struct {
	bun.BaseModel `bun:"table:user"`
	*user.User

	UserAuth []*UserAuth `bun:"rel:has-many,join:id=user_id"`
	UserTag  []*UserTag  `bun:"rel:has-many,join:id=user_id"`
}

func (x *User) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if x.User == nil {
		x.User = &user.User{}
	}

	t := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		if x.CreatedAt == nil {
			x.CreatedAt = t
		}
		if x.UpdatedAt == nil {
			x.UpdatedAt = t
		}
	case *bun.UpdateQuery:
		x.UpdatedAt = t
	}
	return nil
}

func (x *UserAuth) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if x.UserAuth == nil {
		x.UserAuth = &user.UserAuth{}
	}

	t := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		if x.CreatedAt == nil {
			x.CreatedAt = t
		}
		if x.UpdatedAt == nil {
			x.UpdatedAt = t
		}
	case *bun.UpdateQuery:
		x.UpdatedAt = t
	}
	return nil
}

func (*User) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	if _, err := query.DB().NewCreateIndex().IfNotExists().Model((*User)(nil)).
		Index("idx_user").Column("username", "disable", "data", "info").
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (*UserAuth) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	if _, err := query.DB().NewCreateIndex().IfNotExists().Model((*UserAuth)(nil)).
		Index("idx_user_auth").Column("data").
		Exec(ctx); err != nil {
		return err
	}

	if err := foreignKeyIfNotExists(ctx, query.DB(), "user_auth", "user_id", "fk_rel_ua_user_id", "user", "id"); err != nil {
		return err
	}

	return nil
}

func (*UserTag) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	if _, err := query.DB().NewCreateIndex().IfNotExists().Model((*UserTag)(nil)).
		Index("idx_user_tag").Column("source_id", "data").
		Exec(ctx); err != nil {
		return err
	}

	if err := foreignKeyIfNotExists(ctx, query.DB(), "user_tag", "user_id", "fk_rel_user_id", "user", "id"); err != nil {
		return err
	}

	return nil
}

func foreignKeyIfNotExists(ctx context.Context, db *bun.DB,
	table, key string, constraintName string,
	refTable, refKey string) error {

	if has, err := db.NewSelect().Table("information_schema.table_constraints").Where(
		"constraint_type = 'FOREIGN KEY' AND table_name = ? AND constraint_name = ?",
		table, constraintName,
	).Exists(ctx); err != nil || has {
		return err
	}
	if _, err := db.NewRaw(`ALTER TABLE ? ADD CONSTRAINT ? FOREIGN KEY (?) REFERENCES ? (?)`,
		bun.Name(table), bun.Name(constraintName), bun.Name(key), bun.Name(refTable), bun.Name(refKey),
	).Exec(ctx); err != nil {
		return err
	}

	return nil
}
