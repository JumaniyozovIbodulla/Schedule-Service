package service

import (
	"context"
	sa "user/genproto/user_service/super_admins"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (s *SuperAdminService) Create(ctx context.Context, req *sa.CreateSuperAdmin) (*sa.SuperAdmin, error) {
	s.log.Info("Create Super Admin: ", logger.Any("req", req))
	resp, err := s.strg.SuperAdmin().Create(ctx, req)

	if err != nil {
		s.log.Error("Create Super Admin: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *SuperAdminService) Delete(ctx context.Context, req *sa.SuperAdminPrimaryKey) (*sa.Empty, error) {
	s.log.Info("Delete Super Admin:", logger.Any("req", req))

	resp, err := s.strg.SuperAdmin().Delete(ctx, req)

	if err != nil {
		s.log.Error("Delete Super Admin:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *SuperAdminService) GetAll(ctx context.Context, req *sa.GetListSuperAdminRequest) (*sa.GetListSuperAdminResponse, error) {
	s.log.Info("Get all super admins:", logger.Any("req:", req))

	resp, err := s.strg.SuperAdmin().GetAll(ctx, req)

	if err != nil {
		s.log.Error("Get all super admin:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *SuperAdminService) GetById(ctx context.Context, req *sa.SuperAdminPrimaryKey) (*sa.SuperAdmin, error) {
	s.log.Info("Get by id super admin:", logger.Any("req:", req))

	resp, err := s.strg.SuperAdmin().GetById(ctx, req)

	if err != nil {
		s.log.Error("Get by id admin:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *SuperAdminService) Update(ctx context.Context, req *sa.UpdateSuperAdmin) (*sa.SuperAdmin, error) {
	s.log.Info("Update super admin:", logger.Any("req:", req))

	resp, err := s.strg.SuperAdmin().Update(ctx, req)

	if err != nil {
		s.log.Error("Update super admin:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}