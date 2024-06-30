package postgres

import (
	"context"
	ts "schedule/genproto/schedule_service/tasks"
	"schedule/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type taskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) storage.TaskRepo {
	return &taskRepo{
		db: db,
	}
}

func (t *taskRepo) Create(ctx context.Context, req *ts.CreateTask) (*ts.Task, error) {
	resp := &ts.Task{}
	id := uuid.New()

	_, err := t.db.Exec(ctx, `
	INSERT INTO 
		tasks(id, lesson_id, label, deadline, score)
	VALUES($1, $2, $3, $4, $5);`, id, req.LessonId, req.Label, req.Deadline, req.Score)

	if err != nil {
		return nil, err
	}

	resp, err = t.GetById(ctx, &ts.TaskPrimaryKey{Id: id.String()})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *taskRepo) GetById(ctx context.Context, req *ts.TaskPrimaryKey) (*ts.Task, error) {
	resp := &ts.Task{}

	row := t.db.QueryRow(ctx, `
	SELECT 
		id, 
		lesson_id, 
		label, 
		TO_CHAR(deadline,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS deadline,
		score,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		tasks 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(
		&resp.Id, 
		&resp.LessonId, 
		&resp.Label, 
		&resp.Deadline, 
		&resp.Score, 
		&resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *taskRepo) GetAll(ctx context.Context, req *ts.GetListTaskRequest) (*ts.GetListTaskResponse, error) {
	resp := &ts.GetListTaskResponse{}

	rows, err := t.db.Query(ctx, `
	SELECT 
		id, 
		lesson_id, 
		label, 
		deadline, 
		score,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at 
	FROM 
		tasks 
	OFFSET 
		$1 
	LIMIT 
		$2;`, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task ts.Task

		if err = rows.Scan(
			&task.Id,
			&task.LessonId,
			&task.Label,
			&task.Deadline,
			&task.Score,
			&task.CreatedAt); err != nil {
			return nil, err
		}

		resp.Tasks = append(resp.Tasks, &task)
	}

	err = t.db.QueryRow(ctx, `SELECT COUNT(*) FROM tasks;`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *taskRepo) Update(ctx context.Context, req *ts.UpdateTask) (*ts.Task, error) {
	resp := &ts.Task{}
	_, err := t.db.Exec(ctx, `
	UPDATE 
		tasks
	SET
		lesson_id = $2, 
		label = $3, 
		deadline = $4, 
		score = $5
	WHERE
		id = $1;`, req.Id, req.LessonId, req.Label, req.Deadline, req.Score)

	if err != nil {
		return nil, err
	}
	resp, err = t.GetById(ctx, &ts.TaskPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *taskRepo) Delete(ctx context.Context, req *ts.TaskPrimaryKey) (*ts.Empty, error) {
	_, err := t.db.Exec(ctx, `
	DELETE FROM 
		tasks
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}