package service

import (
	"context"

	sch "schedule/genproto/schedule_service/schedules"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (s *ScheduleService) Create(ctx context.Context, schedule *sch.CreateSchedule) (*sch.Schedule, error) {
	s.log.Info("create a schedule:", logger.Any("req:", schedule))
	resp, err := s.strg.Schedule().Create(ctx, schedule)

	if err != nil {
		s.log.Error("failed to create a schedule:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *ScheduleService) Delete(ctx context.Context, schedule *sch.SchedulePrimaryKey) (*sch.Empty, error) {
	s.log.Info("delete a schedule:", logger.Any("req:", schedule))

	_, err := s.strg.Schedule().Delete(ctx, schedule)

	if err != nil {
		s.log.Error("failet to delete a schedule:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (s *ScheduleService) GetAll(ctx context.Context, schedule *sch.GetListScheduleRequest) (*sch.GetListScheduleResponse, error) {
	s.log.Info("get all schedules:", logger.Any("req:", schedule))

	resp, err := s.strg.Schedule().GetAll(ctx, schedule)

	if err != nil {
		s.log.Error("failed to get all schedules:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *ScheduleService) GetById(ctx context.Context, schedule *sch.SchedulePrimaryKey) (*sch.Schedule, error) {
	s.log.Info("get by id schedule:", logger.Any("req:", schedule))

	resp, err := s.strg.Schedule().GetById(ctx, schedule)

	if err != nil {
		s.log.Error("failed to get by id schedule: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *ScheduleService) Update(ctx context.Context, schedule *sch.UpdateSchedule) (*sch.Schedule, error) {
	s.log.Info("Update a schedule:", logger.Any("req:", schedule))

	resp, err := s.strg.Schedule().Update(ctx, schedule)

	if err != nil {
		s.log.Error("failed to get by id schedule:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
