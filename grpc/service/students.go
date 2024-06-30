package service

import (
	"context"
	st "user/genproto/user_service/students"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (s *StudentService) Create(ctx context.Context, student *st.CreateStudent) (*st.Student, error) {
	s.log.Info("create a student:", logger.Any("req:", student))
	resp, err := s.strg.Student().Create(ctx, student)

	if err != nil {
		s.log.Error("failed to create a student:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *StudentService) Delete(ctx context.Context, student *st.StudentPrimaryKey) (*st.Empty, error) {
	s.log.Info("delete a student:", logger.Any("req:", student))

	_, err := s.strg.Student().Delete(ctx, student)

	if err != nil {
		s.log.Error("failet to delete a student:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (s *StudentService) GetAll(ctx context.Context, student *st.GetListStudentsRequest) (*st.GetListStudentsResponse, error) {
	s.log.Info("get all students:", logger.Any("req:", student))

	resp, err := s.strg.Student().GetAll(ctx, student)

	if err != nil {
		s.log.Error("failed to get all students:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *StudentService) GetById(ctx context.Context, student *st.StudentPrimaryKey) (*st.Student, error) {
	s.log.Info("get by id student:", logger.Any("req:", student))

	resp, err := s.strg.Student().GetById(ctx, student)

	if err != nil {
		s.log.Error("failed to get by id student: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *StudentService) Update(ctx context.Context, student *st.UpdateStudent) (*st.Student, error) {
	s.log.Info("Update a student:", logger.Any("req:", student))

	resp, err := s.strg.Student().Update(ctx, student)

	if err != nil {
		s.log.Error("failed to get by id student:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}