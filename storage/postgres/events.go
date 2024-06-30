package postgres

import (
	"context"
	ev "user/genproto/user_service/events"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type eventRepo struct {
	db *pgxpool.Pool
}

func NewEventRepo(db *pgxpool.Pool) storage.EventRepo {
	return &eventRepo{
		db: db,
	}
}

func (e *eventRepo) Create(ctx context.Context, req *ev.CreateEvent) (*ev.Event, error) {
	id := uuid.New()

	_, err := e.db.Exec(ctx, `
		INSERT INTO 
			events(id, topic, descrioption, day, start_time, duration_hours, branch_id)
		VALUES($1, $2, $3, $4, $5, $6, $7);`, id, req.Topic, req.Descrioption, req.Day, req.StartTime, req.DurationHours, req.BranchId)
	if err != nil {
		return nil, err
	}
	event, err := e.GetById(ctx, &ev.EventPrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e *eventRepo) GetById(ctx context.Context, req *ev.EventPrimaryKey) (*ev.Event, error) {
	resp := &ev.Event{}

	row := e.db.QueryRow(ctx, `
	SELECT 
		id, 
		topic, 
		descrioption, 
		day, 
		TO_CHAR(start_time,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS start_time,
		duration_hours,
		branch_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		events 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(
		&resp.Id,
		&resp.Topic,
		&resp.Descrioption,
		&resp.Day,
		&resp.StartTime,
		&resp.DurationHours,
		&resp.BranchId,
		&resp.IsActive,
		&resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *eventRepo) GetAll(ctx context.Context, req *ev.GetListEventsRequest) (*ev.GetListEventsResponse, error) {
	resp := &ev.GetListEventsResponse{}

	rows, err := e.db.Query(ctx, `
	SELECT 
		id, 
		topic, 
		descrioption, 
		day, 
		start_time,
		duration_hours,
		branch_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		events 
	WHERE 
		is_active = 0  
	OFFSET 
		$1 
	LIMIT 
		$2;`, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event ev.Event

		if err = rows.Scan(
			&event.Id,
			&event.Topic,
			&event.Descrioption,
			&event.Day,
			&event.StartTime,
			&event.DurationHours,
			&event.BranchId,
			&event.IsActive,
			&event.CreatedAt); err != nil {
			return nil, err
		}

		resp.Events = append(resp.Events, &event)
	}

	err = e.db.QueryRow(ctx, `SELECT COUNT(*) FROM events WHERE is_active = 0`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *eventRepo) Update(ctx context.Context, req *ev.UpdateEvent) (*ev.Event, error) {
	resp := &ev.Event{}

	_, err := e.db.Exec(ctx, `
	UPDATE 
		events
	SET
		topic = $2, 
		descrioption = $3, 
		day = $4, 
		start_time = $5,
		duration_hours = $6,
		branch_id = $7
	WHERE
		id = $1;`, req.Id, req.Topic, req.Descrioption, req.Day, req.StartTime, req.DurationHours, req.BranchId)

	if err != nil {
		return nil, err
	}
	resp, err = e.GetById(ctx, &ev.EventPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *eventRepo) Delete(ctx context.Context, req *ev.EventPrimaryKey) (*ev.Empty, error) {
	_, err := e.db.Exec(ctx, `
	UPDATE 
		events 
	SET 
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
