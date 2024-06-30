package postgres

import (
	"context"
	ej "user/genproto/user_service/join_events"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type joinEventRepo struct {
	db *pgxpool.Pool
}

func NewEventJoinRepo(db *pgxpool.Pool) storage.EventJoinRepo {
	return &joinEventRepo{
		db: db,
	}
}

func (e *joinEventRepo) Create(ctx context.Context, req *ej.CreateJoinEvent) (*ej.JoinEvent, error) {
	id := uuid.New()

	_, err := e.db.Exec(ctx, `
		INSERT INTO 
			join_events(id, event_id, student_id)
		VALUES($1, $2, $3);`, id, req.EventId, req.StudentId)
	if err != nil {
		return nil, err
	}
	event, err := e.GetById(ctx, &ej.JoinEventPrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e *joinEventRepo) GetById(ctx context.Context, req *ej.JoinEventPrimaryKey) (*ej.JoinEvent, error) {
	resp := &ej.JoinEvent{}

	row := e.db.QueryRow(ctx, `
	SELECT 
		id, 
		event_id, 
		student_id, 
		is_active,
		TO_CHAR(joined_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS joined_at
	FROM 
		join_events 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(
		&resp.Id,
		&resp.EventId,
		&resp.StudentId,
		&resp.IsActive,
		&resp.JoinedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *joinEventRepo) Delete(ctx context.Context, req *ej.JoinEventPrimaryKey) (*ej.Empty, error) {
	_, err := e.db.Exec(ctx, `
	UPDATE 
		join_events 
	SET 
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
