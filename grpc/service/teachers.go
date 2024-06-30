package service

import (
	"context"
	tr "user/genproto/user_service/teachers"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (t *TeacherService) Create(ctx context.Context, teacher *tr.CreateTeacher) (*tr.Teacher, error) {
	t.log.Info("create a teacher:", logger.Any("req:", teacher))
	resp, err := t.strg.Teacher().Create(ctx, teacher)

	if err != nil {
		t.log.Error("failed to create a teacher:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (t *TeacherService) Delete(ctx context.Context, teacher *tr.TeacherPrimaryKey) (*tr.Empty, error) {
	t.log.Info("delete a teacher:", logger.Any("req:", teacher))

	_, err := t.strg.Teacher().Delete(ctx, teacher)

	if err != nil {
		t.log.Error("failet to delete a teacher:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (t *TeacherService) GetAll(ctx context.Context, teacher *tr.GetListTeachersRequest) (*tr.GetListTeachersResponse, error) {
	t.log.Info("get all teachers:", logger.Any("req:", teacher))

	resp, err := t.strg.Teacher().GetAll(ctx, teacher)

	if err != nil {
		t.log.Error("failed to get all teachers:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (t *TeacherService) GetById(ctx context.Context, teacher *tr.TeacherPrimaryKey) (*tr.Teacher, error) {
	t.log.Info("get by id teacher:", logger.Any("req:", teacher))

	resp, err := t.strg.Teacher().GetById(ctx, teacher)

	if err != nil {
		t.log.Error("failed to get by id teacher: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (t *TeacherService) Update(ctx context.Context, teacher *tr.UpdateTeacher) (*tr.Teacher, error) {
	t.log.Info("Update a teacher:", logger.Any("req:", teacher))

	resp, err := t.strg.Teacher().Update(ctx, teacher)

	if err != nil {
		t.log.Error("failed to get by id teacher:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
