package main

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	v1 "github.com/core-pb/user/user/v1"
	"github.com/core-pb/user/user/v1/userconnect"
	"github.com/uptrace/bun"
)

type base struct {
	userconnect.UnimplementedBaseHandler
}

func (base) ListAuthenticate(ctx context.Context, req *connect.Request[v1.ListAuthenticateRequest]) (*connect.Response[v1.ListAuthenticateResponse], error) {
	sq := db.NewSelect().Model(&Authenticate{})
	sq = InOrEqPure(sq, `"authenticate".id`, req.Msg.Id)
	sq = InOrEqPure(sq, `"authenticate".type`, req.Msg.Type)
	sq = QueryFormStruct(sq, `"authenticate".data`, req.Msg.Data)
	sq = QueryFormStruct(sq, `"authenticate".info`, req.Msg.Info)
	if req.Msg.Disable != nil {
		sq = sq.Where(`"authenticate".disable = ?`, *req.Msg.Disable)
	}

	sq = Pagination(sq, req.Msg.Pagination)
	sq = Sorts(sq, req.Msg.Sort)

	var (
		arr        []*v1.Authenticate
		count, err = sq.ScanAndCount(ctx, &arr)
	)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.ListAuthenticateResponse{Data: arr, Count: int64(count)}), nil
}

func (base) AddAuthenticate(ctx context.Context, req *connect.Request[v1.AddAuthenticateRequest]) (*connect.Response[v1.AddAuthenticateResponse], error) {
	val := &Authenticate{Authenticate: &v1.Authenticate{
		Type:    req.Msg.Type,
		Scope:   req.Msg.Scope,
		Disable: req.Msg.Disable,
		Data:    req.Msg.Data,
		Info:    req.Msg.Info,
	}}
	// TODO: check type data

	if _, err := db.NewInsert().Model(val).Returning("*").Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}
	return connect.NewResponse(&v1.AddAuthenticateResponse{Data: val.Authenticate}), nil
}

func (base) SetAuthenticate(ctx context.Context, req *connect.Request[v1.SetAuthenticateRequest]) (*connect.Response[v1.SetAuthenticateResponse], error) {
	val := &Authenticate{}

	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if req.Msg.Scope != nil {
			if has, err := tx.NewSelect().Model(&UserAuth{}).Where("auth_id = ?", req.Msg.Id).Exists(ctx); err != nil {
				return err
			} else if has {
				return errors.New("authenticate is using")
			}
		}

		ts := tx.NewUpdate().Model(val).Where("id IN =", req.Msg.Id).Returning("*")
		if req.Msg.Scope != nil {
			ts.Set("scope = ?", req.Msg.Scope)
		}
		if req.Msg.Disable != nil {
			ts.Set("disable = ?", req.Msg.Disable)
		}
		if req.Msg.Type != nil {
			ts.Set(`"type" = ?`, req.Msg.Type)
		}
		if req.Msg.Data != nil {
			ts.Set(`"data" = ?`, req.Msg.Data)
		}
		if req.Msg.Info != nil {
			ts.Set(`"info" = ?`, req.Msg.Info)
		}

		_, err := ts.Exec(ctx)
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetAuthenticateResponse{Data: val.Authenticate}), nil
}

func (base) DeleteAuthenticate(ctx context.Context, req *connect.Request[v1.DeleteAuthenticateRequest]) (*connect.Response[v1.DeleteAuthenticateResponse], error) {
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&UserAuth{}).Where("auth_id IN (?)", req.Msg.Id).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("authenticate is using")
		}

		_, err := tx.NewDelete().Model(&Authenticate{}).Where("id IN (?)", req.Msg.Id).Exec(ctx)
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DeleteAuthenticateResponse{}), nil
}
