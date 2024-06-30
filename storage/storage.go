package storage

import (
	"context"
	sch "schedule/genproto/schedule_service/schedules"
	ls "schedule/genproto/schedule_service/lessons"
	ts "schedule/genproto/schedule_service/tasks"
	at "schedule/genproto/schedule_service/attendances"
)

type IStorage interface {
	CloseDB()
	Schedule() ScheduleRepo
	Lesson() LessonRepo
	Task() TaskRepo
	Attendance() AttendanceRepo
}

type ScheduleRepo interface {
	Create(ctx context.Context, req *sch.CreateSchedule) (*sch.Schedule, error)
	GetById(ctx context.Context, req *sch.SchedulePrimaryKey) (*sch.Schedule, error)
	GetAll(ctx context.Context, req *sch.GetListScheduleRequest) (*sch.GetListScheduleResponse, error)
	Update(ctx context.Context, req *sch.UpdateSchedule) (*sch.Schedule, error)
	Delete(ctx context.Context, req *sch.SchedulePrimaryKey) (*sch.Empty, error)
}

type LessonRepo interface {
	Create(ctx context.Context, req *ls.CreateLesson) (*ls.Lesson, error)
	GetById(ctx context.Context, req *ls.LessonPrimaryKey) (*ls.Lesson, error)
	GetAll(ctx context.Context, req *ls.GetListLessonRequest) (*ls.GetListLesssonResponse, error)
	Update(ctx context.Context, req *ls.UpdateLesson) (*ls.Lesson, error)
}

type TaskRepo interface {
	Create(ctx context.Context, req *ts.CreateTask) (*ts.Task, error)
	GetById(ctx context.Context, req *ts.TaskPrimaryKey) (*ts.Task, error)
	GetAll(ctx context.Context, req *ts.GetListTaskRequest) (*ts.GetListTaskResponse, error)
	Update(ctx context.Context, req *ts.UpdateTask) (*ts.Task, error)
	Delete(ctx context.Context, req *ts.TaskPrimaryKey) (*ts.Empty, error)
}

type AttendanceRepo interface {
	Create(ctx context.Context, req *at.CreateAttendance) (*at.Attendance, error)
	GetById(ctx context.Context, req *at.AttendancePrimaryKey) (*at.Attendance, error)
	GetAll(ctx context.Context, req *at.GetListAttendanceRequest) (*at.GetListAttendanceResponse, error)
	Update(ctx context.Context, req *at.UpdateAttendance) (*at.Attendance, error)
	Delete(ctx context.Context, req *at.AttendancePrimaryKey) (*at.Empty, error)
}