package grpc

import (
	"user/config"
	ad "user/genproto/user_service/administrators"
	br "user/genproto/user_service/branches"
	ev "user/genproto/user_service/events"
	gr "user/genproto/user_service/groups"
	ej "user/genproto/user_service/join_events"
	mn "user/genproto/user_service/managers"
	st "user/genproto/user_service/students"
	sa "user/genproto/user_service/super_admins"
	sp "user/genproto/user_service/support_teachers"
	ts "user/genproto/user_service/teachers"
	"user/grpc/client"
	"user/grpc/service"
	"user/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.IStorage, superAdmin client.SuperAdminManager, branch client.BrancnesManager, groups client.GroupsManager, manager client.ManagersManager, admins client.AdminsManager, supportTeacher client.SupportTeachersManager, teacher client.TeachersManager, student client.StudentsManager, events client.EventsManager, joinEvent client.JoinEventsManager) *grpc.Server {
	grpcServer := grpc.NewServer()

	sa.RegisterSuperAdminServiceServer(grpcServer, service.NewSuperAdminService(cfg, log, strg, superAdmin))
	br.RegisterBranchesServiceServer(grpcServer, service.NewBranchesService(cfg, log, strg, branch))
	gr.RegisterGroupsServiceServer(grpcServer, service.NewGroupsService(cfg, log, strg, groups))
	mn.RegisterManagersServiceServer(grpcServer, service.NewManagersService(cfg, log, strg, manager))
	ad.RegisterAdministratorsServiceServer(grpcServer, service.NewAdminsService(cfg, log, strg, admins))
	sp.RegisterSupportTeacherServiceServer(grpcServer, service.NewSupportTeacherService(cfg, log, strg, supportTeacher))
	ts.RegisterTeacherServiceServer(grpcServer, service.NewTeachersService(cfg, log, strg, teacher))
	st.RegisterStudentServiceServer(grpcServer, service.NewStudentsService(cfg, log, strg, student))
	ev.RegisterEventServiceServer(grpcServer, service.NewEventsService(cfg, log, strg, events))
	ej.RegisterJoinEventServiceServer(grpcServer, service.NewJoinEventsService(cfg, log, strg, joinEvent))
	
	reflection.Register(grpcServer)
	return grpcServer
}
