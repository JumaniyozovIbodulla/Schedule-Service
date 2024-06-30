package service

import (
	"context"

	at "schedule/genproto/schedule_service/attendances"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (a *AttendanceService) Create(ctx context.Context, attendance *at.CreateAttendance) (*at.Attendance, error) {
	a.log.Info("create a attendance:", logger.Any("req:", attendance))
	resp, err := a.strg.Attendance().Create(ctx, attendance)

	if err != nil {
		a.log.Error("failed to create a attendance:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AttendanceService) GetById(ctx context.Context, attendance *at.AttendancePrimaryKey) (*at.Attendance, error) {
	a.log.Info("get by id attendance:", logger.Any("req:", attendance))

	resp, err := a.strg.Attendance().GetById(ctx, attendance)

	if err != nil {
		a.log.Error("failed to get by id attendance: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AttendanceService) Delete(ctx context.Context, attendance *at.AttendancePrimaryKey) (*at.Empty, error) {
	a.log.Info("delete a attendance:", logger.Any("req:", attendance))

	_, err := a.strg.Attendance().Delete(ctx, attendance)

	if err != nil {
		a.log.Error("failet to delete a attendance:", logger.Error(err))
		return nil, err
	}
	return nil, nil
}

func (a *AttendanceService) GetAll(ctx context.Context, attendance *at.GetListAttendanceRequest) (*at.GetListAttendanceResponse, error) {
	a.log.Info("get all attendances:", logger.Any("req:", attendance))

	resp, err := a.strg.Attendance().GetAll(ctx, attendance)

	if err != nil {
		a.log.Error("failed to get all attendances:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AttendanceService) Update(ctx context.Context, attendance *at.UpdateAttendance) (*at.Attendance, error) {
	a.log.Info("Update a lesson:", logger.Any("req:", attendance))

	resp, err := a.strg.Attendance().Update(ctx, attendance)

	if err != nil {
		a.log.Error("failed to update attendance:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}