package client

import "schedule/config"

type ScheduleManager interface{}

type LessonManager interface{}

type TaskManager interface{}

type AttendanceManager interface{}


type grpcScheduleClients struct{}

type grpcLessonClients struct{}

type grpcTaskClients struct{}

type grpcAttendanceClients struct{}


func NewGrpcScheduleClients(cfg config.Config) (ScheduleManager, error) {
	return *&grpcScheduleClients{}, nil
}

func NewGrpcLessonClients(cfg config.Config) (LessonManager, error) {
	return *&grpcLessonClients{}, nil
}

func NewGrpcTaskClients(cfg config.Config) (TaskManager, error) {
	return *&grpcTaskClients{}, nil
}

func NewGrpcAttendanceClients(cfg config.Config) (AttendanceManager, error) {
	return *&grpcAttendanceClients{}, nil
}