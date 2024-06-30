package service

import (
	"context"
	mn "user/genproto/user_service/managers"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (m *ManagersService) Create(ctx context.Context, manager *mn.CreateManager) (*mn.Manager, error) {
	m.log.Info("create a manager:", logger.Any("req:", manager))
	resp, err := m.strg.Manager().Create(ctx, manager)

	if err != nil {
		m.log.Error("failed to create a manager:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (m *ManagersService) Delete(ctx context.Context, manager *mn.ManagerPrimaryKey) (*mn.Empty, error) {
	m.log.Info("delete a manager:", logger.Any("req:", manager))

	_, err := m.strg.Manager().Delete(ctx, manager)

	if err != nil {
		m.log.Error("failet to delete a manager:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (m *ManagersService) GetAll(ctx context.Context, manager *mn.GetListManagersRequest) (*mn.GetListManagersResponse, error) {
	m.log.Info("get all managers:", logger.Any("req:", manager))

	resp, err := m.strg.Manager().GetAll(ctx, manager)

	if err != nil {
		m.log.Error("failed to get all managers:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (m *ManagersService) GetById(ctx context.Context, manager *mn.ManagerPrimaryKey) (*mn.Manager, error) {
	m.log.Info("get by id manager:", logger.Any("req:", manager))

	resp, err := m.strg.Manager().GetById(ctx, manager)

	if err != nil {
		m.log.Error("failed to get by id manager: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (m *ManagersService) Update(ctx context.Context, manager *mn.UpdateManager) (*mn.Manager, error) {
	m.log.Info("Update a manager:", logger.Any("req:", manager))

	resp, err := m.strg.Manager().Update(ctx, manager)

	if err != nil {
		m.log.Error("failed to get by id manager:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
