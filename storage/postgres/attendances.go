package postgres

import (
	"context"
	at "schedule/genproto/schedule_service/attendances"
	"schedule/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

func mapAttendToPostgreSQL(attend at.AttendType) string {
	switch attend {
	case at.AttendType_attended:
		return "attended"
	case at.AttendType_absent:
		return "absent"
	case at.AttendType_late:
		return "late"
	default:
		return ""
	}
}

func mapPostgreSQLToAttendType(attend string) at.AttendType {
	switch attend {
	case "attended":
		return at.AttendType_attended
	case "absent":
		return at.AttendType_absent
	case "late":
		return at.AttendType_late
	default:
		return at.AttendType(0)
	}
}

type attendanceRepo struct {
	db *pgxpool.Pool
}

func NewAttendanceRepo(db *pgxpool.Pool) storage.AttendanceRepo {
	return &attendanceRepo{
		db: db,
	}
}

func (a *attendanceRepo) Create(ctx context.Context, req *at.CreateAttendance) (*at.Attendance, error) {
	id := uuid.New()

	status := mapAttendToPostgreSQL(req.Status)

	_, err := a.db.Exec(ctx, `
		INSERT INTO 
			attendances(id, lesson_id, student_id, status, late_minute)
		VALUES($1, $2, $3, $4, $5);`, id, req.LessonId, req.StudentId, status, req.LateMinute)
	if err != nil {
		return nil, err
	}

	attendance, err := a.GetById(ctx, &at.AttendancePrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (a *attendanceRepo) Delete(ctx context.Context, req *at.AttendancePrimaryKey) (*at.Empty, error) {
	_, err := a.db.Exec(ctx, `
	DELETE FROM 
		attendances 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (a *attendanceRepo) GetAll(ctx context.Context, req *at.GetListAttendanceRequest) (*at.GetListAttendanceResponse, error) {
	resp := &at.GetListAttendanceResponse{}

	rows, err := a.db.Query(ctx, ` 
	SELECT 
		id,
		lesson_id,
		student_id, 
		status,
		late_minute
	FROM
		attendances
	OFFSET
		$1
	LIMIT
		$2;`, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			attendance at.Attendance
			status     pgtype.Text
		)

		if err = rows.Scan(
			&attendance.Id,
			&attendance.LessonId,
			&attendance.StudentId,
			&status,
			&attendance.LateMinute); err != nil {
			return nil, err
		}

		attendance.Status = mapPostgreSQLToAttendType(status.String)

		resp.Attendances = append(resp.Attendances, &attendance)
	}
	err = a.db.QueryRow(ctx, `SELECT COUNT(*) FROM attendances;`).Scan(&resp.Count)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *attendanceRepo) GetById(ctx context.Context, req *at.AttendancePrimaryKey) (*at.Attendance, error) {

	var (
		status     pgtype.Text
		attendance at.Attendance
	)

	row := a.db.QueryRow(ctx, `
	SELECT
		id,
		lesson_id,
		student_id, 
		status,
		late_minute
	FROM
		attendances
	WHERE
		id = $1;`, req.Id)

	err := row.Scan(
		&attendance.Id,
		&attendance.LessonId,
		&attendance.StudentId,
		&status,
		&attendance.LateMinute)

	if err != nil {
		return nil, err
	}

	attendance.Status = mapPostgreSQLToAttendType(status.String)

	return &attendance, nil
}

func (a *attendanceRepo) Update(ctx context.Context, req *at.UpdateAttendance) (*at.Attendance, error) {
	status := mapAttendToPostgreSQL(req.Status)

	_, err := a.db.Exec(ctx, `
	UPDATE 
		attendances
	SET
		lesson_id = $2,
		student_id = $3,
		status = $4,
		late_minute = $5
	WHERE
		id = $1;`,
		req.Id,
		req.LessonId,
		req.StudentId,
		status,
		req.LateMinute)

	if err != nil {
		return nil, err
	}

	resp, err := a.GetById(ctx, &at.AttendancePrimaryKey{Id: req.Id})

	if err != nil {
		return nil, err
	}
	return resp, nil
}
