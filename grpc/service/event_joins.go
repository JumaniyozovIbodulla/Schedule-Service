package service

import (
	"context"
	ej "user/genproto/user_service/join_events"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (e *JoinEventsService) Create(ctx context.Context, eventJoin *ej.CreateJoinEvent) (*ej.JoinEvent, error) {
	e.log.Info("create a eventJoin:", logger.Any("req:", eventJoin))
	resp, err := e.strg.EventJoin().Create(ctx, eventJoin)

	if err != nil {
		e.log.Error("failed to create a eventJoin:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (e *JoinEventsService) GetById(ctx context.Context, eventJoin *ej.JoinEventPrimaryKey) (*ej.JoinEvent, error) {
	e.log.Info("get by id eventJoin:", logger.Any("req:", eventJoin))

	resp, err := e.strg.EventJoin().GetById(ctx, eventJoin)

	if err != nil {
		e.log.Error("failed to get by id eventJoin: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (e *JoinEventsService) Delete(ctx context.Context, eventJoin *ej.JoinEventPrimaryKey) (*ej.Empty, error) {
	e.log.Info("delete a eventJoin:", logger.Any("req:", eventJoin))

	_, err := e.strg.EventJoin().Delete(ctx, eventJoin)

	if err != nil {
		e.log.Error("failet to delete a eventJoin:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}