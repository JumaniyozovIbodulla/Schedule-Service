package service

import (
	"context"
	br "user/genproto/user_service/branches"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (b *BranchesService) Create(ctx context.Context, branch *br.CreateBranch) (*br.Branch, error) {
	b.log.Info("create branch:", logger.Any("req:", branch))
	resp, err := b.strg.Branch().Create(ctx, branch)

	if err != nil {
		b.log.Error("failed to create a branch:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (b *BranchesService) Delete(ctx context.Context, branch *br.BranchePrimaryKey) (*br.Empty, error) {
	b.log.Info("delete a branch:", logger.Any("req:", branch))

	_, err := b.strg.Branch().Delete(ctx, branch)

	if err != nil {
		b.log.Error("failet to delete a branch:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (b *BranchesService) GetAll(ctx context.Context, branch *br.GetListBranchesRequest) (*br.GetListBranchesResponse, error) {
	b.log.Info("get all branches:", logger.Any("req:", branch))

	resp, err := b.strg.Branch().GetAll(ctx, branch)

	if err != nil {
		b.log.Error("failed to get all branches:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (b *BranchesService) GetById(ctx context.Context, branch *br.BranchePrimaryKey) (*br.Branch, error) {
	b.log.Info("get by id branch:", logger.Any("req:", branch))

	resp, err := b.strg.Branch().GetById(ctx, branch)

	if err != nil {
		b.log.Error("failed to get by id branch: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (b *BranchesService) Update(ctx context.Context, branch *br.UpdateBranch) (*br.Branch, error) {
	b.log.Info("Update a branch:", logger.Any("req:", branch))

	resp, err := b.strg.Branch().Update(ctx, branch)

	if err != nil {
		b.log.Error("failed to get by id branch:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
