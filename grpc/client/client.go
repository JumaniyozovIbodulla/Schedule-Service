package client

import "user/config"

type SuperAdminManager interface{}

type BrancnesManager interface{}

type GroupsManager interface{}

type ManagersManager interface{}

type AdminsManager interface{}

type SupportTeachersManager interface{}

type TeachersManager interface{}

type StudentsManager interface{}

type EventsManager interface{}

type JoinEventsManager interface{}


type grpcSuperAdminClients struct{}

type grpcBranchesClients struct{}

type grpcGroupsClients struct{}

type grpcManagersClients struct{}

type grpcAdminsClients struct{}

type grpcSupportteacherClients struct{}

type grpcTeachersClients struct{}

type grpcStudentClients struct{}

type grpcEventClients struct{}

type grpcEventJoinClients struct{}


func NewGrpcSuperAdminsClients(cfg config.Config) (SuperAdminManager, error) {
	return *&grpcSuperAdminClients{}, nil
}

func NewGrpcBranchesClients(cfg config.Config) (BrancnesManager, error) {
	return *&grpcBranchesClients{}, nil
}

func NewGrpcGroupsClients(cfg config.Config) (GroupsManager, error) {
	return *&grpcGroupsClients{}, nil
}


func NewGrpcManagersClients(cfg config.Config) (ManagersManager, error) {
	return *&grpcManagersClients{}, nil
}

func NewGrpcAdminsClients(cfg config.Config) (AdminsManager, error) {
	return *&grpcAdminsClients{}, nil
}

func NewGrpcSupportTeachersClients(cfg config.Config) (SupportTeachersManager, error) {
	return *&grpcSupportteacherClients{}, nil
}

func NewGrpcTeachersClients(cfg config.Config) (TeachersManager, error) {
	return *&grpcTeachersClients{}, nil
}

func NewGrpcStudentsClients(cfg config.Config) (TeachersManager, error) {
	return *&grpcTeachersClients{}, nil
}

func NewGrpcEventsClients(cfg config.Config) (EventsManager, error) {
	return *&grpcEventClients{}, nil
}

func NewGrpcEventJoinsClients(cfg config.Config) (JoinEventsManager, error) {
	return *&grpcEventJoinClients{}, nil
}