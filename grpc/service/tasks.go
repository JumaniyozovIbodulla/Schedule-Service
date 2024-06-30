package service

import (
	"context"
	ts "schedule/genproto/schedule_service/tasks"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (t *TasksService) Create(ctx context.Context, task *ts.CreateTask) (*ts.Task, error) {
	t.log.Info("create task:", logger.Any("req:", task))
	resp, err := t.strg.Task().Create(ctx, task)

	if err != nil {
		t.log.Error("failed to create a task:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (t *TasksService) Delete(ctx context.Context, task *ts.TaskPrimaryKey) (*ts.Empty, error) {
	t.log.Info("delete a task:", logger.Any("req:", task))

	_, err := t.strg.Task().Delete(ctx, task)

	if err != nil {
		t.log.Error("failet to delete a task:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (t *TasksService) GetAll(ctx context.Context, task *ts.GetListTaskRequest) (*ts.GetListTaskResponse, error) {
	t.log.Info("get all tasks:", logger.Any("req:", task))

	resp, err := t.strg.Task().GetAll(ctx, task)

	if err != nil {
		t.log.Error("failed to get all tasks:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (t *TasksService) GetById(ctx context.Context, task *ts.TaskPrimaryKey) (*ts.Task, error) {
	t.log.Info("get by id task:", logger.Any("req:", task))

	resp, err := t.strg.Task().GetById(ctx, task)

	if err != nil {
		t.log.Error("failed to get by id task: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (t *TasksService) Update(ctx context.Context, task *ts.UpdateTask) (*ts.Task, error) {
	t.log.Info("Update a task:", logger.Any("req:", task))

	resp, err := t.strg.Task().Update(ctx, task)

	if err != nil {
		t.log.Error("failed to update task:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
