package grpc

import (
	"schedule/config"
	sch "schedule/genproto/schedule_service/schedules"
	ts "schedule/genproto/schedule_service/tasks"
	ls "schedule/genproto/schedule_service/lessons"
	at "schedule/genproto/schedule_service/attendances"
	"schedule/grpc/client"
	"schedule/grpc/service"
	"schedule/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.IStorage, schedule client.ScheduleManager, lesson client.LessonManager, tasks client.TaskManager, attendance client.AttendanceManager) *grpc.Server {
	grpcServer := grpc.NewServer()

	sch.RegisterScheduleServiceServer(grpcServer, service.NewScheduleService(cfg, log, strg, schedule))
	ts.RegisterTaskServiceServer(grpcServer, service.NewTasksService(cfg, log, strg, tasks))
	ls.RegisterLesssonServiceServer(grpcServer, service.NewLessonsService(cfg, log, strg, lesson))
	at.RegisterAttendanceServiceServer(grpcServer, service.NewAttendanceService(cfg, log, strg, attendance))

	reflection.Register(grpcServer)
	return grpcServer
}