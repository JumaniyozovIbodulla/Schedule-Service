package main

import (
	"context"
	"net"
	"user/config"
	"user/grpc"
	"user/grpc/client"
	"user/storage/postgres"

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

	superAdmin, err := client.NewGrpcSuperAdminsClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcSuperAdminsClients: ", logger.Error(err))
	}

	branch, err := client.NewGrpcBranchesClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcBranchesClients: ", logger.Error(err))
	}

	groups, err := client.NewGrpcGroupsClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcGroupsClients: ", logger.Error(err))
	}

	managers, err := client.NewGrpcManagersClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcManagersClients: ", logger.Error(err))
	}

	admins, err := client.NewGrpcAdminsClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcAdminsClients: ", logger.Error(err))
	}

	supportTeacher, err := client.NewGrpcSupportTeachersClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcSupportTeachersClients: ", logger.Error(err))
	}

	teacher, err := client.NewGrpcTeachersClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcTeachersClients: ", logger.Error(err))
	}

	student, err := client.NewGrpcStudentsClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcStudentsClients: ", logger.Error(err))
	}

	events, err := client.NewGrpcEventsClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcEventsClients: ", logger.Error(err))
	}

	joinevents, err := client.NewGrpcEventJoinsClients(cfg)

	if err != nil {
		log.Panic("client.NewGrpcEventJoinsClients: ", logger.Error(err))
	}

	grpcServer := grpc.SetUpServer(cfg, log, pgStore, superAdmin, branch, groups, managers, admins, supportTeacher, teacher, student, events, joinevents)

	lis, err := net.Listen("tcp", cfg.ContentGRPCPort)

	if err != nil {
		log.Panic("net.Listen: ", logger.Error(err))
	}

	log.Info("GRPC: Server being started...", logger.String("port", cfg.ContentGRPCPort))

	if err := grpcServer.Serve(lis); err != nil {
		log.Panic("grpcServer.Serve: ", logger.Error(err))
	}
}