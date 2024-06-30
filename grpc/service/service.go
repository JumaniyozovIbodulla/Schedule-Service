package service

import (
	"user/config"
	"user/grpc/client"
	"user/storage"
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
	"github.com/saidamir98/udevs_pkg/logger"
)

type SuperAdminService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.SuperAdminManager
	*sa.UnimplementedSuperAdminServiceServer
}

type BranchesService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.BrancnesManager
	*br.UnimplementedBranchesServiceServer
}

type GroupsService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.GroupsManager
	*gr.UnimplementedGroupsServiceServer
}

type ManagersService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.ManagersManager
	*mn.UnimplementedManagersServiceServer
}

type AdminsService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.AdminsManager
	*ad.UnimplementedAdministratorsServiceServer
}

type SupportTeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.SupportTeachersManager
	*sp.UnimplementedSupportTeacherServiceServer
}

type TeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.TeachersManager
	*ts.UnimplementedTeacherServiceServer
}

type StudentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.StudentsManager
	*st.UnimplementedStudentServiceServer
}

type EventsService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.EventsManager
	*ev.UnimplementedEventServiceServer
}

type JoinEventsService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.JoinEventsManager
	*ej.UnimplementedJoinEventServiceServer
}

func NewSuperAdminService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.SuperAdminManager) *SuperAdminService {
	return &SuperAdminService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewBranchesService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.BrancnesManager) *BranchesService {
	return &BranchesService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewGroupsService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.GroupsManager) *GroupsService {
	return &GroupsService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewManagersService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.ManagersManager) *ManagersService {
	return &ManagersService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewAdminsService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.AdminsManager) *AdminsService {
	return &AdminsService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewSupportTeacherService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.SuperAdminManager) *SupportTeacherService {
	return &SupportTeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewTeachersService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.TeachersManager) *TeacherService {
	return &TeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewStudentsService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.StudentsManager) *StudentService {
	return &StudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewEventsService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.EventsManager) *EventsService {
	return &EventsService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func NewJoinEventsService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvc client.JoinEventsManager) *JoinEventsService {
	return &JoinEventsService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}
