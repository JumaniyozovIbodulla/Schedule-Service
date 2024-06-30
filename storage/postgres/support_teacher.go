package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	sp "user/genproto/user_service/support_teachers"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type supportTeacherRepo struct {
	db *pgxpool.Pool
}

func NewSupportTeacher(db *pgxpool.Pool) storage.SupportTeacherRepo {
	return &supportTeacherRepo{
		db: db,
	}
}

func (s *supportTeacherRepo) Create(ctx context.Context, req *sp.CreateSupportTeacher) (*sp.SupportTeacher, error) {
	resp := &sp.SupportTeacher{}
	id := uuid.New()

	var externalId string

	row := s.db.QueryRow(ctx, `SELECT extra_id FROM support_teachers ORDER BY created_at DESC LIMIT 1;`)

	err := row.Scan(&externalId)
	if err != nil {
		_, err = s.db.Exec(ctx, `
			INSERT INTO 
				support_teachers(id, extra_id, fullname, phone, password, ielts_score, ielts_attepms_count, salary, months_worked, branch_id, group_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`, id, "ST00001", req.FullName, req.Phone, req.Password, req.IeltsScore, req.IeltsAttepmsCount, req.Salary, req.MonthsWorked, req.BranchId, req.GroupId)
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
		result := fmt.Sprintf("ST%s%d", zeros, number)
		_, err = s.db.Exec(ctx, `
			INSERT INTO 
				support_teachers(id, extra_id, fullname, phone, password, ielts_score, ielts_attepms_count, salary, months_worked, branch_id, group_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`, id, result, req.FullName, req.Phone, req.Password, req.IeltsScore, req.IeltsAttepmsCount, req.Salary, req.MonthsWorked, req.BranchId, req.GroupId)

		if err != nil {
			return nil, err
		}
	}

	resp, err = s.GetById(ctx, &sp.SupportTeacherPrimaryKey{Id: id.String()})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *supportTeacherRepo) GetById(ctx context.Context, req *sp.SupportTeacherPrimaryKey) (*sp.SupportTeacher, error) {
	resp := &sp.SupportTeacher{}

	row := s.db.QueryRow(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		ielts_score, 
		ielts_attepms_count,
		salary, 
		months_worked,
		branch_id,
		group_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		support_teachers 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(
		&resp.Id,
		&resp.ExtraId,
		&resp.FullName,
		&resp.Phone,
		&resp.Password,
		&resp.IeltsScore,
		&resp.IeltsAttepmsCount,
		&resp.Salary,
		&resp.MonthsWorked,
		&resp.BranchId,
		&resp.GroupId,
		&resp.IsActive,
		&resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *supportTeacherRepo) GetAll(ctx context.Context, req *sp.GetListSupportTeachersRequest) (*sp.GetListSupportTeachersResponse, error) {
	resp := &sp.GetListSupportTeachersResponse{}

	rows, err := s.db.Query(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		ielts_score, 
		ielts_attepms_count,
		salary, 
		months_worked,
		branch_id,
		group_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		support_teachers 
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
		var supportTeacher sp.SupportTeacher

		if err = rows.Scan(
			&supportTeacher.Id,
			&supportTeacher.ExtraId,
			&supportTeacher.FullName,
			&supportTeacher.Phone,
			&supportTeacher.Password,
			&supportTeacher.IeltsScore,
			&supportTeacher.IeltsAttepmsCount,
			&supportTeacher.Salary,
			&supportTeacher.MonthsWorked,
			&supportTeacher.BranchId,
			&supportTeacher.GroupId,
			&supportTeacher.IsActive,
			&supportTeacher.CreatedAt); err != nil {
			return nil, err
		}

		resp.SupportTeachers = append(resp.SupportTeachers, &supportTeacher)
	}

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) FROM support_teachers WHERE is_active = 0`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *supportTeacherRepo) Update(ctx context.Context, req *sp.UpdateSupportTeacher) (*sp.SupportTeacher, error) {
	resp := &sp.SupportTeacher{}

	_, err := s.db.Exec(ctx, `
	UPDATE 
		support_teachers
	SET
		fullname = $2, 
		phone = $3, 
		password = $4, 
		ielts_score = $5,
		ielts_attepms_count = $6,
		salary = $7, 
		months_worked = $8,
		branch_id = $9,
		group_id = $10
	WHERE
		id = $1;`, req.Id, req.FullName, req.Phone, req.Password, req.IeltsScore, req.IeltsAttepmsCount, req.Salary, req.MonthsWorked, req.BranchId, req.GroupId)

	if err != nil {
		return nil, err
	}
	resp, err = s.GetById(ctx, &sp.SupportTeacherPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *supportTeacherRepo) Delete(ctx context.Context, req *sp.SupportTeacherPrimaryKey) (*sp.Empty, error) {
	_, err := s.db.Exec(ctx, `
	UPDATE 
		support_teachers 
	SET 
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}