package postgres

import (
	"context"
	ls "schedule/genproto/schedule_service/lessons"
	"schedule/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type lessonRepo struct {
	db *pgxpool.Pool
}

func NewLessonRepo(db *pgxpool.Pool) storage.LessonRepo {
	return &lessonRepo{
		db: db,
	}
}

func (l *lessonRepo) Create(ctx context.Context, req *ls.CreateLesson) (*ls.Lesson, error) {
	id := uuid.New()

	_, err := l.db.Exec(ctx, `
		INSERT INTO 
			lessons(id, schedule_id)
		VALUES($1, $2);`, id, req.ScheduleId)
	if err != nil {
		return nil, err
	}
	lesson, err := l.GetById(ctx, &ls.LessonPrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}
	return lesson, nil
}

func (l* lessonRepo) GetById(ctx context.Context, req *ls.LessonPrimaryKey) (*ls.Lesson, error) {
	resp := &ls.Lesson{}

	row := l.db.QueryRow(ctx, `
	SELECT 
		id, 
		schedule_id,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		lessons 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(
		&resp.Id,
		&resp.ScheduleId,
		&resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}


func (l *lessonRepo) GetAll(ctx context.Context, req *ls.GetListLessonRequest) (*ls.GetListLesssonResponse, error) {
	resp := &ls.GetListLesssonResponse{}

	rows, err := l.db.Query(ctx, `
	SELECT 
		id, 
		schedule_id,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		lessons  
	OFFSET 
		$1 
	LIMIT 
		$2;`, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var lesson ls.Lesson

		if err = rows.Scan(
			&lesson.Id,
			&lesson.ScheduleId,
			&lesson.CreatedAt); err != nil {
			return nil, err
		}

		resp.Lessons = append(resp.Lessons, &lesson)
	}

	err = l.db.QueryRow(ctx, `SELECT COUNT(*) FROM lessons;`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *lessonRepo) Update(ctx context.Context, req *ls.UpdateLesson) (*ls.Lesson, error) {
	resp := &ls.Lesson{}

	_, err := l.db.Exec(ctx, `
	UPDATE 
		lessons
	SET
		schedule_id = $2
	WHERE
		id = $1;`, req.Id, req.ScheduleId)

	if err != nil {
		return nil, err
	}

	resp, err = l.GetById(ctx, &ls.LessonPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}