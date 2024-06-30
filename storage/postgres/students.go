package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	st "user/genproto/user_service/students"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type studentRepo struct {
	db *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) storage.StudentRepo {
	return &studentRepo{
		db: db,
	}
}

func (s *studentRepo) Create(ctx context.Context, req *st.CreateStudent) (*st.Student, error) {
	resp := &st.Student{}
	id := uuid.New()

	var externalId string

	row := s.db.QueryRow(ctx, `SELECT extra_id FROM students ORDER BY created_at DESC LIMIT 1;`)

	err := row.Scan(&externalId)
	if err != nil {
		_, err = s.db.Exec(ctx, `
			INSERT INTO 
				students(id, extra_id, fullname, phone, password, paid_sum, course_count, total_sum, branch_id, group_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`, id, "S00001", req.FullName, req.Phone, req.Password, req.PaidSum, req.CourseCount, req.TotalSum, req.BranchId, req.GroupId)
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
		result := fmt.Sprintf("S%s%d", zeros, number)
		_, err = s.db.Exec(ctx, `
			INSERT INTO 
				students(id, extra_id, fullname, phone, password, paid_sum, course_count, total_sum, branch_id, group_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`, id, result, req.FullName, req.Phone, req.Password, req.PaidSum, req.CourseCount, req.TotalSum, req.BranchId, req.GroupId)

		if err != nil {
			return nil, err
		}
	}

	resp, err = s.GetById(ctx, &st.StudentPrimaryKey{Id: id.String()})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *studentRepo) GetById(ctx context.Context, req *st.StudentPrimaryKey) (*st.Student, error) {
	resp := &st.Student{}

	row := s.db.QueryRow(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		paid_sum,
		course_count,
		total_sum, 
		branch_id,
		group_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		students 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(
		&resp.Id,
		&resp.ExtraId,
		&resp.FullName,
		&resp.Phone,
		&resp.Password,
		&resp.PaidSum,
		&resp.CourseCount,
		&resp.TotalSum,
		&resp.BranchId,
		&resp.GroupId,
		&resp.IsActive,
		&resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *studentRepo) GetAll(ctx context.Context, req *st.GetListStudentsRequest) (*st.GetListStudentsResponse, error) {
	resp := &st.GetListStudentsResponse{}

	rows, err := s.db.Query(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		paid_sum,
		course_count,
		total_sum, 
		branch_id,
		group_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		students 
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
		var student st.Student

		if err = rows.Scan(
			&student.Id,
			&student.ExtraId,
			&student.FullName,
			&student.Phone,
			&student.Password,
			&student.PaidSum,
			&student.CourseCount,
			&student.TotalSum,
			&student.BranchId,
			&student.GroupId,
			&student.IsActive,
			&student.CreatedAt); err != nil {
			return nil, err
		}

		resp.Students = append(resp.Students, &student)
	}

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) FROM students WHERE is_active = 0`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *studentRepo) Update(ctx context.Context, req *st.UpdateStudent) (*st.Student, error) {
	resp := &st.Student{}

	_, err := s.db.Exec(ctx, `
	UPDATE 
		students
	SET
		fullname = $2, 
		phone = $3, 
		password = $4, 
		paid_sum = $5,
		course_count = $6,
		total_sum = $7, 
		branch_id = $8,
		group_id = $9
	WHERE
		id = $1;`, req.Id, req.FullName, req.Phone, req.Password, req.PaidSum, req.CourseCount, req.TotalSum, req.BranchId, req.GroupId)

	if err != nil {
		return nil, err
	}
	resp, err = s.GetById(ctx, &st.StudentPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *studentRepo) Delete(ctx context.Context, req *st.StudentPrimaryKey) (*st.Empty, error) {
	_, err := s.db.Exec(ctx, `
	UPDATE 
		students 
	SET 
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
