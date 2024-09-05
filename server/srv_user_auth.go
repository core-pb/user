package main

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/core-pb/authenticate/authenticate/v1"
	v1 "github.com/core-pb/user/user/v1"
	"github.com/uptrace/bun"
)

func (base) SetUserAuth(ctx context.Context, req *connect.Request[v1.SetUserAuthRequest]) (*connect.Response[v1.SetUserAuthResponse], error) {
	resp, err := auth.Generate(ctx, connect.NewRequest(&authenticate.GenerateRequest{
		Id:   req.Msg.AuthId,
		Data: req.Msg.Data,
	}))
	if err != nil {
		return nil, err
	}

	ua := &UserAuth{UserAuth: &v1.UserAuth{
		UserId: req.Msg.UserId,
		AuthId: req.Msg.AuthId,
		Data:   resp.Msg.Data,
		Info:   req.Msg.Info,
	}}

	if _, err = db.NewInsert().Model(ua).On("CONFLICT (user_id,auth_id) DO UPDATE").
		Set("scope = EXCLUDED.scope, data = EXCLUDED.data, info = EXCLUDED.info").
		Returning("*").Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetUserAuthResponse{Data: ua.UserAuth}), nil
}

func (base) DeleteUserAuth(ctx context.Context, req *connect.Request[v1.DeleteUserAuthRequest]) (*connect.Response[v1.DeleteUserAuthResponse], error) {
	if len(req.Msg.UserId) == 0 && len(req.Msg.AuthId) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid argument"))
	}
	tx := db.NewDelete().Model(&UserAuth{})
	if len(req.Msg.UserId) != 0 {
		tx = tx.Where("user_id IN (?)", bun.In(req.Msg.UserId))
	}
	if len(req.Msg.AuthId) != 0 {
		tx = tx.Where("auth_id IN (?)", bun.In(req.Msg.AuthId))
	}

	if _, err := tx.Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}
	return connect.NewResponse(&v1.DeleteUserAuthResponse{}), nil
}
