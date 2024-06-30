package service

import (
	"context"
	ls "schedule/genproto/schedule_service/lessons"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (l *LessonsService) Create(ctx context.Context, lesson *ls.CreateLesson) (*ls.Lesson, error) {
	l.log.Info("create a lesson:", logger.Any("req:", lesson))
	resp, err := l.strg.Lesson().Create(ctx, lesson)

	if err != nil {
		l.log.Error("failed to create a lesson:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (l *LessonsService) GetAll(ctx context.Context, lesson *ls.GetListLessonRequest) (*ls.GetListLesssonResponse, error) {
	l.log.Info("get all lessons:", logger.Any("req:", lesson))

	resp, err := l.strg.Lesson().GetAll(ctx, lesson)

	if err != nil {
		l.log.Error("failed to get all lessons:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (l *LessonsService) GetById(ctx context.Context, lesson *ls.LessonPrimaryKey) (*ls.Lesson, error) {
	l.log.Info("get by id lesson:", logger.Any("req:", lesson))

	resp, err := l.strg.Lesson().GetById(ctx, lesson)

	if err != nil {
		l.log.Error("failed to get by id lesson: ", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (l *LessonsService) Update(ctx context.Context, lesson *ls.UpdateLesson) (*ls.Lesson, error) {
	l.log.Info("Update a lesson:", logger.Any("req:", lesson))

	resp, err := l.strg.Lesson().Update(ctx, lesson)

	if err != nil {
		l.log.Error("failed to get by id lesson:", logger.Error(err))
		return nil, err
	}
	return resp, nil
}