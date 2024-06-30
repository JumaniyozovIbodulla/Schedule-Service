package main

import (
	"context"
	"net"
	"schedule/config"
	"schedule/grpc"
	"schedule/grpc/client"
	"schedule/storage/postgres"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	cfg := config.Load()

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)

	defer logger.Cleanup(log)

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)

	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}

	defer pgStore.CloseDB()

	schedule, err := client.NewGrpcScheduleClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcScheduleClients: ", logger.Error(err))
	}

	lesson, err := client.NewGrpcLessonClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcLessonClients: ", logger.Error(err))
	}

	tasks, err := client.NewGrpcTaskClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcTaskClients: ", logger.Error(err))
	}

	attendance, err := client.NewGrpcAttendanceClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcAttendanceClients: ", logger.Error(err))
	}

	grpcServer := grpc.SetUpServer(cfg, log, pgStore, schedule, lesson, tasks, attendance)

	lis, err := net.Listen("tcp", cfg.ContentGRPCPort)

	if err != nil {
		log.Panic("net.Listen: ", logger.Error(err))
	}

	log.Info("GRPC: Server being started...", logger.String("port", cfg.ContentGRPCPort))

	if err := grpcServer.Serve(lis); err != nil {
		log.Panic("grpcServer.Serve: ", logger.Error(err))
	}
}