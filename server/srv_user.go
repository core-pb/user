package main

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/core-pb/user/user/v1"
	"github.com/core-pb/user/user/v1/userconnect"
	"github.com/uptrace/bun"
)

type base struct {
	userconnect.UnimplementedBaseHandler
}

func (base) ListUser(ctx context.Context, req *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error) {
	sq := db.NewSelect().Model(&User{})
	sq = InOrEqPure(sq, `"user".id`, req.Msg.Id)
	sq = InOrEqPure(sq, `"user".username`, req.Msg.Username)
	sq = QueryFormStruct(sq, `"user".data`, req.Msg.Data)
	sq = QueryFormStruct(sq, `"user".info`, req.Msg.Info)
	if req.Msg.Disable != nil {
		sq = sq.Where(`"user".disable = ?`, *req.Msg.Disable)
	}

	if req.Msg.TagId != nil {
		sq = sq.Join(`LEFT JOIN "user_tag" ON "user_tag".user_id = "user".id`)
		sq = InOrEqPure(sq, `"user_tag".tag_id`, req.Msg.TagId)
	}

	sq = Pagination(sq, req.Msg.Pagination)
	sq = Sorts(sq, req.Msg.Sort)

	var (
		arr        []*UserDetail
		ids        []uint64
		count, err = sq.ScanAndCount(ctx, &ids)
	)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	if err = db.NewSelect().
		Model(&arr).
		Relation("UserInfo").
		Relation("UserTag").
		Where(`"id" IN (?)`, bun.In(ids)).
		Scan(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.ListUserResponse{
		Data: Array2Array(arr, func(a1 *UserDetail) *v1.UserDetail {
			return &v1.UserDetail{
				User:     a1.User,
				UserAuth: Array2Array(a1.UserAuth, func(a1 *UserAuth) *v1.UserAuth { return a1.UserAuth }),
				UserTag:  Array2Array(a1.UserTag, func(a1 *UserTag) *v1.UserTag { return a1.UserTag }),
			}
		}), Count: int64(count),
	}), nil
}

func (base) AddUser(ctx context.Context, req *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error) {
	val := &User{User: &v1.User{
		Username: req.Msg.Username,
		Disable:  req.Msg.Disable,
		Data:     req.Msg.Data,
		Info:     req.Msg.Info,
	}}

	if _, err := db.NewInsert().Model(val).Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.AddUserResponse{Data: val.User}), nil
}

func (base) SetUser(ctx context.Context, req *connect.Request[v1.SetUserRequest]) (*connect.Response[v1.SetUserResponse], error) {
	val := new(User)

	tx := db.NewUpdate().Model(val).Returning("*").Where("id = ?", req.Msg.Id)
	if req.Msg.Disable != nil {
		tx = tx.Set("disable = ?", *req.Msg.Disable)
	}
	if req.Msg.Data != nil {
		tx = tx.Set(`"data" = ?`, req.Msg.Data)
	}

	if _, err := tx.Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetUserResponse{Data: val.User}), nil
}

func (base) SetUserInfo(ctx context.Context, req *connect.Request[v1.SetUserInfoRequest]) (*connect.Response[v1.SetUserInfoResponse], error) {
	val := new(User)

	if _, err := db.NewUpdate().Model(val).Returning("*").Where("id = ?", req.Msg.Id).Set("info = ?", req.Msg.Info).Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetUserInfoResponse{Data: val.User}), nil
}

func (base) DeleteUser(ctx context.Context, req *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	if _, err := db.NewDelete().Model(&User{}).Where("id IN (?)", bun.In(req.Msg.Id)).Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DeleteUserResponse{}), nil
}
