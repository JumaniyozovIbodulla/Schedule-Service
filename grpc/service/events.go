package service

import (
	"context"
	ev "user/genproto/user_service/events"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (e *EventsService) Create(ctx context.Context, event *ev.CreateEvent) (*ev.Event, error) {
	e.log.Info("create a event:", logger.Any("req:", event))
	resp, err := e.strg.Event().Create(ctx, event)

	if err != nil {
		e.log.Error("failed to create a event:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (e *EventsService) Delete(ctx context.Context, event *ev.EventPrimaryKey) (*ev.Empty, error) {
	e.log.Info("delete a event:", logger.Any("req:", event))

	_, err := e.strg.Event().Delete(ctx, event)

	if err != nil {
		e.log.Error("failet to delete a event:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (e *EventsService) GetAll(ctx context.Context, event *ev.GetListEventsRequest) (*ev.GetListEventsResponse, error) {
	e.log.Info("get all events:", logger.Any("req:", event))

	resp, err := e.strg.Event().GetAll(ctx, event)

	if err != nil {
		e.log.Error("failed to get all events:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (e *EventsService) GetById(ctx context.Context, event *ev.EventPrimaryKey) (*ev.Event, error) {
	e.log.Info("get by id event:", logger.Any("req:", event))

	resp, err := e.strg.Event().GetById(ctx, event)

	if err != nil {
		e.log.Error("failed to get by id event: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (e *EventsService) Update(ctx context.Context, event *ev.UpdateEvent) (*ev.Event, error) {
	e.log.Info("Update a event:", logger.Any("req:", event))

	resp, err := e.strg.Event().Update(ctx, event)

	if err != nil {
		e.log.Error("failed to get by id event:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
