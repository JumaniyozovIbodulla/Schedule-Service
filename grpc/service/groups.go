package service

import (
	"context"
	gr "user/genproto/user_service/groups"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (g *GroupsService) Create(ctx context.Context, group *gr.CreateGroup) (*gr.Group, error) {
	g.log.Info("create group:", logger.Any("req:", group))
	resp, err := g.strg.Groups().Create(ctx, group)

	if err != nil {
		g.log.Error("failed to create a group:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (g *GroupsService) Delete(ctx context.Context, group *gr.GroupPrimaryKey) (*gr.Empty, error) {
	g.log.Info("delete a group:", logger.Any("req:", group))

	_, err := g.strg.Groups().Delete(ctx, group)

	if err != nil {
		g.log.Error("failet to delete a group:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (g *GroupsService) GetAll(ctx context.Context, group *gr.GetListGroupsRequest) (*gr.GetListGroupsResponse, error) {
	g.log.Info("get all groups:", logger.Any("req:", group))

	resp, err := g.strg.Groups().GetAll(ctx, group)

	if err != nil {
		g.log.Error("failed to get all groups:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (g *GroupsService) GetById(ctx context.Context, group *gr.GroupPrimaryKey) (*gr.Group, error) {
	g.log.Info("get by id group:", logger.Any("req:", group))

	resp, err := g.strg.Groups().GetById(ctx, group)

	if err != nil {
		g.log.Error("failed to get by id group: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (g *GroupsService) Update(ctx context.Context, group *gr.UpdateGroup) (*gr.Group, error) {
	g.log.Info("Update a group:", logger.Any("req:", group))

	resp, err := g.strg.Groups().Update(ctx, group)

	if err != nil {
		g.log.Error("failed to get by id group:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
