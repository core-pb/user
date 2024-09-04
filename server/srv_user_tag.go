package main

import (
	"context"

	"connectrpc.com/connect"
	"github.com/core-pb/tag/client"
	"github.com/core-pb/tag/tag/v1"
	v1 "github.com/core-pb/user/user/v1"
	"github.com/core-pb/user/user/v1/userconnect"
	"github.com/uptrace/bun"
)

type _tag struct {
	userconnect.UnimplementedTagHandler
}

func (_tag) GetTag(ctx context.Context, _ *connect.Request[v1.GetTagRequest]) (*connect.Response[v1.GetTagResponse], error) {
	resp, err := client.Internal().GetAllByModule(ctx, connect.NewRequest(&tag.GetAllByModuleRequest{ModuleId: module.Id}))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&v1.GetTagResponse{Type: resp.Msg.Type, Tag: resp.Msg.Tag}), nil
}

func (_tag) SetTagType(ctx context.Context, req *connect.Request[v1.SetTagTypeRequest]) (*connect.Response[v1.SetTagTypeResponse], error) {
	if _, err := client.Internal().SetTypeWithModule(ctx, connect.NewRequest(&tag.SetTypeWithModuleRequest{
		ModuleId:  module.Id,
		TypeId:    req.Msg.Id,
		Info:      req.Msg.Info,
		Inherit:   req.Msg.Inherit,
		Exclusive: req.Msg.Exclusive,
	})); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.SetTagTypeResponse{}), nil
}

func (_tag) DeleteTagType(ctx context.Context, req *connect.Request[v1.DeleteTagTypeRequest]) (*connect.Response[v1.DeleteTagTypeResponse], error) {
	if _, err := client.Internal().DeleteTypeWithModule(ctx, connect.NewRequest(&tag.DeleteTypeWithModuleRequest{
		ModuleId: module.Id,
		TypeId:   req.Msg.Id,
	})); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DeleteTagTypeResponse{}), nil
}

func (_tag) SetTag(ctx context.Context, req *connect.Request[v1.SetTagRequest]) (*connect.Response[v1.SetTagResponse], error) {
	if _, err := client.Internal().SetTagWithModule(ctx, connect.NewRequest(&tag.SetTagWithModuleRequest{
		ModuleId: module.Id,
		TagId:    req.Msg.Id,
		TypeId:   req.Msg.TypeId,
		ParentId: req.Msg.ParentId,
		Data:     req.Msg.Data,
		Info:     req.Msg.Info,
	})); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.SetTagResponse{}), nil
}

func (_tag) DeleteTag(ctx context.Context, req *connect.Request[v1.DeleteTagRequest]) (*connect.Response[v1.DeleteTagResponse], error) {
	if _, err := client.Internal().DeleteTagWithModule(ctx, connect.NewRequest(&tag.DeleteTagWithModuleRequest{
		ModuleId: module.Id,
		TagId:    req.Msg.Id,
	})); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DeleteTagResponse{}), nil
}

func (_tag) SetUserTag(ctx context.Context, req *connect.Request[v1.SetUserTagRequest]) (*connect.Response[v1.SetUserTagResponse], error) {
	resp, err := client.Internal().BindRelation(ctx, connect.NewRequest(&tag.BindRelationRequest{
		ModuleId:   module.Id,
		ExternalId: req.Msg.UserId,
		TagId:      req.Msg.TagId,
		Data:       req.Msg.Data,
	}))
	if err != nil {
		return nil, err
	}

	if err = db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if len(resp.Msg.CleanTagId) != 0 {
			if _, err = tx.NewSelect().Model(&UserTag{}).Where(`user_id = ? AND tag_id IN (?)`, req.Msg.UserId, bun.In(resp.Msg.CleanTagId)).Exec(ctx); err != nil {
				return err
			}
		}

		arr := make([]*UserTag, 0, 1+len(resp.Msg.InheritTagId))
		arr = append(arr, &UserTag{UserTag: &v1.UserTag{
			UserId: req.Msg.UserId,
			TagId:  req.Msg.TagId,
			Data:   req.Msg.Data,
		}})

		for _, v := range resp.Msg.InheritTagId {
			arr = append(arr, &UserTag{UserTag: &v1.UserTag{
				UserId:   req.Msg.UserId,
				TagId:    v,
				SourceId: req.Msg.TagId,
			}})
		}

		if _, err = tx.NewInsert().On("CONFLICT (user_id,tag_id) DO UPDATE").
			Set("source_id = EXCLUDED.source_id, data = EXCLUDED.data").Model(&arr).Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetUserTagResponse{}), nil
}

func (_tag) DeleteUserTag(ctx context.Context, req *connect.Request[v1.DeleteUserTagRequest]) (*connect.Response[v1.DeleteUserTagResponse], error) {
	resp, err := client.Internal().UnbindRelation(ctx, connect.NewRequest(&tag.UnbindRelationRequest{
		ModuleId:   module.Id,
		ExternalId: req.Msg.UserId,
		TagId:      req.Msg.TagId,
	}))
	if err != nil {
		return nil, err
	}

	if err = db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if len(resp.Msg.CleanTagId) != 0 {
			if _, err = tx.NewSelect().Model(&UserTag{}).Where(`user_id = ? AND tag_id IN (?)`,
				req.Msg.UserId, bun.In(resp.Msg.CleanTagId)).Exec(ctx); err != nil {
				return err
			}
		}

		arr := make([]uint64, 0, 1+len(resp.Msg.CleanTagId))
		arr = append(arr, req.Msg.TagId)
		arr = append(arr, resp.Msg.CleanTagId...)

		if _, err = tx.NewDelete().Model(&UserTag{}).Where(`user_id = ? AND tag_id IN (?)`,
			req.Msg.UserId, bun.In(resp.Msg.CleanTagId)).Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DeleteUserTagResponse{}), nil
}
