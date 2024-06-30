package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	tr "user/genproto/user_service/teachers"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type teacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) storage.TeacherRepo {
	return &teacherRepo{
		db: db,
	}
}

func (t *teacherRepo) Create(ctx context.Context, req *tr.CreateTeacher) (*tr.Teacher, error) {
	resp := &tr.Teacher{}
	id := uuid.New()

	var externalId string

	row := t.db.QueryRow(ctx, `SELECT extra_id FROM teachers ORDER BY created_at DESC LIMIT 1;`)

	err := row.Scan(&externalId)
	if err != nil {
		_, err = t.db.Exec(ctx, `
			INSERT INTO 
				teachers(id, extra_id, fullname, phone, password, salary, months_worked, ielts_score, ielts_attepms_count, support_teacher_id, branch_id, group_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`, id, "T00001", req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore, req.IeltsAttepmsCount, req.SupportTeacherId, req.BranchId, req.GroupId)
		if err != nil {
			return nil, err
		}
	} else {
		takeNumber := externalId[2:]
		number, err := strconv.Atoi(takeNumber)
		if err != nil {
			return nil, err
		}

		number++
		stringNumber := strconv.Itoa(number)
		length := 5 - len(stringNumber)
		zeros := strings.Repeat("0", length)
		result := fmt.Sprintf("T%s%d", zeros, number)
		_, err = t.db.Exec(ctx, `
			INSERT INTO 
				teachers(id, extra_id, fullname, phone, password, salary, months_worked, ielts_score, ielts_attepms_count, support_teacher_id, branch_id, group_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`, id, result, req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore, req.IeltsAttepmsCount, req.SupportTeacherId, req.BranchId, req.GroupId)

		if err != nil {
			return nil, err
		}
	}

	resp, err = t.GetById(ctx, &tr.TeacherPrimaryKey{Id: id.String()})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *teacherRepo) GetById(ctx context.Context, req *tr.TeacherPrimaryKey) (*tr.Teacher, error) {
	resp := &tr.Teacher{}

	row := t.db.QueryRow(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		salary,
		months_worked,
		ielts_score, 
		ielts_attepms_count,
		support_teacher_id,
		branch_id,
		group_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		teachers 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(
		&resp.Id,
		&resp.ExtraId,
		&resp.FullName,
		&resp.Phone,
		&resp.Password,
		&resp.Salary,
		&resp.MonthsWorked,
		&resp.IeltsScore,
		&resp.IeltsAttepmsCount,
		&resp.SupportTeacherId,
		&resp.BranchId,
		&resp.GroupId,
		&resp.IsActive,
		&resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *teacherRepo) GetAll(ctx context.Context, req *tr.GetListTeachersRequest) (*tr.GetListTeachersResponse, error) {
	resp := &tr.GetListTeachersResponse{}

	rows, err := t.db.Query(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		salary,
		months_worked,
		ielts_score, 
		ielts_attepms_count,
		support_teacher_id,
		branch_id,
		group_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		teachers 
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
		var teacher tr.Teacher

		if err = rows.Scan(
			&teacher.Id,
			&teacher.ExtraId,
			&teacher.FullName,
			&teacher.Phone,
			&teacher.Password,
			&teacher.Salary,
			&teacher.MonthsWorked,
			&teacher.IeltsScore,
			&teacher.IeltsAttepmsCount,
			&teacher.SupportTeacherId,
			&teacher.BranchId,
			&teacher.GroupId,
			&teacher.IsActive,
			&teacher.CreatedAt); err != nil {
			return nil, err
		}

		resp.Teachers = append(resp.Teachers, &teacher)
	}

	err = t.db.QueryRow(ctx, `SELECT COUNT(*) FROM teachers WHERE is_active = 0`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *teacherRepo) Update(ctx context.Context, req *tr.UpdateTeacher) (*tr.Teacher, error) {
	resp := &tr.Teacher{}

	_, err := t.db.Exec(ctx, `
	UPDATE 
		teachers
	SET
		fullname = $2, 
		phone = $3, 
		password = $4, 
		salary = $5,
		months_worked = $6,
		ielts_score = $7, 
		ielts_attepms_count = $8,
		support_teacher_id = $9,
		branch_id = $10,
		group_id = $11
	WHERE
		id = $1;`, req.Id, req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore, req.IeltsAttepmsCount, req.SupportTeacherId, req.BranchId, req.GroupId)

	if err != nil {
		return nil, err
	}
	resp, err = t.GetById(ctx, &tr.TeacherPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *teacherRepo) Delete(ctx context.Context, req *tr.TeacherPrimaryKey) (*tr.Empty, error) {
	_, err := t.db.Exec(ctx, `
	UPDATE 
		teachers 
	SET 
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}