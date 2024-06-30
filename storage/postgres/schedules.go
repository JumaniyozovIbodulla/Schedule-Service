package postgres

import (
	"context"
	sch "schedule/genproto/schedule_service/schedules"
	"schedule/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type scheduleRepo struct {
	db *pgxpool.Pool
}

func NewScheduleRepo(db *pgxpool.Pool) storage.ScheduleRepo {
	return &scheduleRepo{
		db: db,
	}
}

func (s *scheduleRepo) Create(ctx context.Context, req *sch.CreateSchedule) (*sch.Schedule, error) {
	id := uuid.New()

	_, err := s.db.Exec(ctx, `
		INSERT INTO 
			schedules(id, group_id, start_time, end_time, branch_id, teacher_id, support_teacher_id)
		VALUES($1, $2, $3, $4, $5, $6, $7);`, id, req.GroupId, req.StartTime, req.EndTime, req.BranchId, req.TeacherId, req.SupportTeacherId)
	if err != nil {
		return nil, err
	}
	schedule, err := s.GetById(ctx, &sch.SchedulePrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *scheduleRepo) GetById(ctx context.Context, req *sch.SchedulePrimaryKey) (*sch.Schedule, error) {
	resp := &sch.Schedule{}

	row := s.db.QueryRow(ctx, `
	SELECT 
		id, 
		group_id, 
		TO_CHAR(start_time,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS start_time, 
		TO_CHAR(end_time,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS end_time, 
		branch_id,
		teacher_id,
		support_teacher_id,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		schedules 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(
		&resp.Id,
		&resp.GroupId,
		&resp.StartTime,
		&resp.EndTime,
		&resp.BranchId,
		&resp.TeacherId,
		&resp.SupportTeacherId,
		&resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *scheduleRepo) GetAll(ctx context.Context, req *sch.GetListScheduleRequest) (*sch.GetListScheduleResponse, error) {
	resp := &sch.GetListScheduleResponse{}

	rows, err := s.db.Query(ctx, `
	SELECT 
		id, 
		group_id, 
		TO_CHAR(start_time,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS start_time, 
		TO_CHAR(end_time,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS end_time, 
		branch_id,
		teacher_id,
		branch_id,
		support_teacher_id,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		schedules  
	OFFSET 
		$1 
	LIMIT 
		$2;`, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schedule sch.Schedule

		if err = rows.Scan(
			&schedule.Id,
			&schedule.GroupId,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.BranchId,
			&schedule.TeacherId,
			&schedule.SupportTeacherId,
			&schedule.CreatedAt); err != nil {
			return nil, err
		}

		resp.Schedules = append(resp.Schedules, &schedule)
	}

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) FROM schedules;`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *scheduleRepo) Update(ctx context.Context, req *sch.UpdateSchedule) (*sch.Schedule, error) {
	resp := &sch.Schedule{}

	_, err := s.db.Exec(ctx, `
	UPDATE 
		schedules
	SET
		group_id = $2, 
		start_time = $3::timestamptz, 
		end_time = $4::timestamptz, 
		branch_id = $5,
		teacher_id = $6,
		support_teacher_id = $7
	WHERE
		id = $1;`, req.Id, req.GroupId, req.StartTime, req.EndTime, req.BranchId, req.TeacherId, req.SupportTeacherId)

	if err != nil {
		return nil, err
	}
	resp, err = s.GetById(ctx, &sch.SchedulePrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *scheduleRepo) Delete(ctx context.Context, req *sch.SchedulePrimaryKey) (*sch.Empty, error) {
	_, err := s.db.Exec(ctx, `
	DELETE FROM 
		schedules 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
