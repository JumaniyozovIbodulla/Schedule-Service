package postgres

import (
	"context"
	"fmt"
	"user/config"
	"user/storage"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db             *pgxpool.Pool
	superAdmin     storage.SuperAdminRepo
	branch         storage.BranchRepo
	group          storage.GroupRepo
	manager        storage.ManagerRepo
	admins         storage.AdminRepo
	supportTeacher storage.SupportTeacherRepo
	teacher        storage.TeacherRepo
	student        storage.StudentRepo
	event          storage.EventRepo
	eventJoin      storage.EventJoinRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2)
	args = append(args, level, msg)

	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}

	fmt.Println(args...)
}

func (s *Store) SuperAdmin() storage.SuperAdminRepo {
	if s.superAdmin == nil {
		s.superAdmin = NewSuperAdminRepo(s.db)
	}
	return s.superAdmin
}

func (s *Store) Branch() storage.BranchRepo {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}
	return s.branch
}

func (s *Store) Groups() storage.GroupRepo {
	if s.group == nil {
		s.group = NewGroupRepo(s.db)
	}
	return s.group
}

func (s *Store) Manager() storage.ManagerRepo {
	if s.manager == nil {
		s.manager = NewManagerRepo(s.db)
	}
	return s.manager
}

func (s *Store) Admins() storage.AdminRepo {
	if s.admins == nil {
		s.admins = NewAdminRepo(s.db)
	}
	return s.admins
}

func (s *Store) SupportTeacher() storage.SupportTeacherRepo {
	if s.supportTeacher == nil {
		s.supportTeacher = NewSupportTeacher(s.db)
	}
	return s.supportTeacher
}

func (s *Store) Teacher() storage.TeacherRepo {
	if s.teacher == nil {
		s.teacher = NewTeacherRepo(s.db)
	}
	return s.teacher
}

func (s *Store) Student() storage.StudentRepo {
	if s.student == nil {
		s.student = NewStudentRepo(s.db)
	}
	return s.student
}

func (s *Store) Event() storage.EventRepo {
	if s.event == nil {
		s.event = NewEventRepo(s.db)
	}
	return s.event
}

func (s *Store) EventJoin() storage.EventJoinRepo {
	if s.eventJoin == nil {
		s.eventJoin = NewEventJoinRepo(s.db)
	}
	return s.eventJoin
}
