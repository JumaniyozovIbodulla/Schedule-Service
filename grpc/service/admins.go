package service

import (
	"context"
	ad "user/genproto/user_service/administrators"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (a *AdminsService) Create(ctx context.Context, admin *ad.CreateAdminstrator) (*ad.Adminstrator, error) {
	a.log.Info("create a admin:", logger.Any("req:", admin))
	resp, err := a.strg.Admins().Create(ctx, admin)

	if err != nil {
		a.log.Error("failed to create a admin:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AdminsService) Delete(ctx context.Context, admin *ad.AdminstratorPrimaryKey) (*ad.Empty, error) {
	a.log.Info("delete a admin:", logger.Any("req:", admin))

	_, err := a.strg.Admins().Delete(ctx, admin)

	if err != nil {
		a.log.Error("failet to delete a admin:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (a *AdminsService) GetAll(ctx context.Context, admin *ad.GetListAdminstratorsRequest) (*ad.GetListAdminstratorsResponse, error) {
	a.log.Info("get all admins:", logger.Any("req:", admin))

	resp, err := a.strg.Admins().GetAll(ctx, admin)

	if err != nil {
		a.log.Error("failed to get all admins:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AdminsService) GetById(ctx context.Context, admin *ad.AdminstratorPrimaryKey) (*ad.Adminstrator, error) {
	a.log.Info("get by id admin:", logger.Any("req:", admin))

	resp, err := a.strg.Admins().GetById(ctx, admin)

	if err != nil {
		a.log.Error("failed to get by id admin: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AdminsService) Update(ctx context.Context, admin *ad.UpdateAdminstrator) (*ad.Adminstrator, error) {
	a.log.Info("Update a admin:", logger.Any("req:", admin))

	resp, err := a.strg.Admins().Update(ctx, admin)

	if err != nil {
		a.log.Error("failed to get by id admin:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
