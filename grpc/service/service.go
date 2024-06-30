package service

import (
	"schedule/config"
	sch "schedule/genproto/schedule_service/schedules"
	ls "schedule/genproto/schedule_service/lessons"
	ts "schedule/genproto/schedule_service/tasks"
	at "schedule/genproto/schedule_service/attendances"
	
	"schedule/grpc/client"
	"schedule/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ScheduleService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.ScheduleManager
	*sch.UnimplementedScheduleServiceServer
}

type LessonsService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.LessonManager
	*ls.UnimplementedLesssonServiceServer
}

type TasksService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.TaskManager
	*ts.UnimplementedTaskServiceServer
}

type AttendanceService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.AttendanceManager
	*at.UnimplementedAttendanceServiceServer
}


func NewScheduleService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.ScheduleManager) *ScheduleService {
	return &ScheduleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewLessonsService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.LessonManager) *LessonsService {
	return &LessonsService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewTasksService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.TaskManager) *TasksService {
	return &TasksService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewAttendanceService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.AttendanceManager) *AttendanceService {
	return &AttendanceService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}