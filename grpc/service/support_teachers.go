package service

import (
	"context"
	sp "user/genproto/user_service/support_teachers"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (s *SupportTeacherService) Create(ctx context.Context, supportTeacher *sp.CreateSupportTeacher) (*sp.SupportTeacher, error) {
	s.log.Info("create a supportTeacher:", logger.Any("req:", supportTeacher))
	resp, err := s.strg.SupportTeacher().Create(ctx, supportTeacher)

	if err != nil {
		s.log.Error("failed to create a supportTeacher:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *SupportTeacherService) Delete(ctx context.Context, supportTeacher *sp.SupportTeacherPrimaryKey) (*sp.Empty, error) {
	s.log.Info("delete a supportTeacher:", logger.Any("req:", supportTeacher))

	_, err := s.strg.SupportTeacher().Delete(ctx, supportTeacher)

	if err != nil {
		s.log.Error("failet to delete a supportTeacher:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (s *SupportTeacherService) GetAll(ctx context.Context, supportTeacher *sp.GetListSupportTeachersRequest) (*sp.GetListSupportTeachersResponse, error) {
	s.log.Info("get all supportTeachers:", logger.Any("req:", supportTeacher))

	resp, err := s.strg.SupportTeacher().GetAll(ctx, supportTeacher)

	if err != nil {
		s.log.Error("failed to get all supportTeachers:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *SupportTeacherService) GetById(ctx context.Context, supportTeacher *sp.SupportTeacherPrimaryKey) (*sp.SupportTeacher, error) {
	s.log.Info("get by id supportTeacher:", logger.Any("req:", supportTeacher))

	resp, err := s.strg.SupportTeacher().GetById(ctx, supportTeacher)

	if err != nil {
		s.log.Error("failed to get by id supportTeacher: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *SupportTeacherService) Update(ctx context.Context, supportTeacher *sp.UpdateSupportTeacher) (*sp.SupportTeacher, error) {
	s.log.Info("Update a supportTeacher:", logger.Any("req:", supportTeacher))

	resp, err := s.strg.SupportTeacher().Update(ctx, supportTeacher)

	if err != nil {
		s.log.Error("failed to get by id supportTeacher:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
