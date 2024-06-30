package storage

import (
	"context"
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
)

type IStorage interface {
	CloseDB()
	SuperAdmin() SuperAdminRepo
	Branch() BranchRepo
	Groups() GroupRepo
	Manager() ManagerRepo
	Admins() AdminRepo
	SupportTeacher() SupportTeacherRepo
	Teacher() TeacherRepo
	Student() StudentRepo
	Event() EventRepo
	EventJoin() EventJoinRepo
}

type SuperAdminRepo interface {
	Create(ctx context.Context, req *sa.CreateSuperAdmin) (*sa.SuperAdmin, error)
	GetById(ctx context.Context, req *sa.SuperAdminPrimaryKey) (*sa.SuperAdmin, error)
	GetAll(ctx context.Context, req *sa.GetListSuperAdminRequest) (*sa.GetListSuperAdminResponse, error)
	Update(ctx context.Context, req *sa.UpdateSuperAdmin) (*sa.SuperAdmin, error)
	Delete(ctx context.Context, req *sa.SuperAdminPrimaryKey) (*sa.Empty, error)
}

type BranchRepo interface {
	Create(ctx context.Context, req *br.CreateBranch) (*br.Branch, error)
	GetById(ctx context.Context, req *br.BranchePrimaryKey) (*br.Branch, error)
	GetAll(ctx context.Context, req *br.GetListBranchesRequest) (*br.GetListBranchesResponse, error)
	Update(ctx context.Context, req *br.UpdateBranch) (*br.Branch, error)
	Delete(ctx context.Context, req *br.BranchePrimaryKey) (*br.Empty, error)
}

type GroupRepo interface {
	Create(ctx context.Context, req *gr.CreateGroup) (*gr.Group, error)
	GetById(ctx context.Context, req *gr.GroupPrimaryKey) (*gr.Group, error)
	GetAll(ctx context.Context, req *gr.GetListGroupsRequest) (*gr.GetListGroupsResponse, error)
	Update(ctx context.Context, req *gr.UpdateGroup) (*gr.Group, error)
	Delete(ctx context.Context, req *gr.GroupPrimaryKey) (*gr.Empty, error)
}

type ManagerRepo interface {
	Create(ctx context.Context, req *mn.CreateManager) (*mn.Manager, error)
	GetById(ctx context.Context, req *mn.ManagerPrimaryKey) (*mn.Manager, error)
	GetAll(ctx context.Context, req *mn.GetListManagersRequest) (*mn.GetListManagersResponse, error)
	Update(ctx context.Context, req *mn.UpdateManager) (*mn.Manager, error)
	Delete(ctx context.Context, req *mn.ManagerPrimaryKey) (*mn.Empty, error)
}

type AdminRepo interface {
	Create(ctx context.Context, req *ad.CreateAdminstrator) (*ad.Adminstrator, error)
	GetById(ctx context.Context, req *ad.AdminstratorPrimaryKey) (*ad.Adminstrator, error)
	GetAll(ctx context.Context, req *ad.GetListAdminstratorsRequest) (*ad.GetListAdminstratorsResponse, error)
	Update(ctx context.Context, req *ad.UpdateAdminstrator) (*ad.Adminstrator, error)
	Delete(ctx context.Context, req *ad.AdminstratorPrimaryKey) (*ad.Empty, error)
}

type SupportTeacherRepo interface {
	Create(ctx context.Context, req *sp.CreateSupportTeacher) (*sp.SupportTeacher, error)
	GetById(ctx context.Context, req *sp.SupportTeacherPrimaryKey) (*sp.SupportTeacher, error)
	GetAll(ctx context.Context, req *sp.GetListSupportTeachersRequest) (*sp.GetListSupportTeachersResponse, error)
	Update(ctx context.Context, req *sp.UpdateSupportTeacher) (*sp.SupportTeacher, error)
	Delete(ctx context.Context, req *sp.SupportTeacherPrimaryKey) (*sp.Empty, error)
}

type TeacherRepo interface {
	Create(ctx context.Context, req *ts.CreateTeacher) (*ts.Teacher, error)
	GetById(ctx context.Context, req *ts.TeacherPrimaryKey) (*ts.Teacher, error)
	GetAll(ctx context.Context, req *ts.GetListTeachersRequest) (*ts.GetListTeachersResponse, error)
	Update(ctx context.Context, req *ts.UpdateTeacher) (*ts.Teacher, error)
	Delete(ctx context.Context, req *ts.TeacherPrimaryKey) (*ts.Empty, error)
}

type StudentRepo interface {
	Create(ctx context.Context, req *st.CreateStudent) (*st.Student, error)
	GetById(ctx context.Context, req *st.StudentPrimaryKey) (*st.Student, error)
	GetAll(ctx context.Context, req *st.GetListStudentsRequest) (*st.GetListStudentsResponse, error)
	Update(ctx context.Context, req *st.UpdateStudent) (*st.Student, error)
	Delete(ctx context.Context, req *st.StudentPrimaryKey) (*st.Empty, error)
}

type EventRepo interface {
	Create(ctx context.Context, req *ev.CreateEvent) (*ev.Event, error)
	GetById(ctx context.Context, req *ev.EventPrimaryKey) (*ev.Event, error)
	GetAll(ctx context.Context, req *ev.GetListEventsRequest) (*ev.GetListEventsResponse, error)
	Update(ctx context.Context, req *ev.UpdateEvent) (*ev.Event, error)
	Delete(ctx context.Context, req *ev.EventPrimaryKey) (*ev.Empty, error)
}

type EventJoinRepo interface {
	Create(ctx context.Context, req *ej.CreateJoinEvent) (*ej.JoinEvent, error)
	GetById(ctx context.Context, req *ej.JoinEventPrimaryKey) (*ej.JoinEvent, error)
	Delete(ctx context.Context, req *ej.JoinEventPrimaryKey) (*ej.Empty, error)
}
