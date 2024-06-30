package postgres

import (
	"context"
	"fmt"
	"schedule/config"
	"schedule/storage"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db         *pgxpool.Pool
	schedule   storage.ScheduleRepo
	lesson     storage.LessonRepo
	task       storage.TaskRepo
	attendance storage.AttendanceRepo
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

func (s *Store) Schedule() storage.ScheduleRepo {
	if s.schedule == nil {
		s.schedule = NewScheduleRepo(s.db)
	}
	return s.schedule
}

func (s *Store) Lesson() storage.LessonRepo {
	if s.lesson == nil {
		s.lesson = NewLessonRepo(s.db)
	}
	return s.lesson
}

func (s *Store) Task() storage.TaskRepo {
	if s.task == nil {
		s.task = NewTaskRepo(s.db)
	}
	return s.task
}

func (s *Store) Attendance() storage.AttendanceRepo {
	if s.attendance == nil {
		s.attendance = NewAttendanceRepo(s.db)
	}
	return s.attendance
}
