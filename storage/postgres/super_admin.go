package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	sp "user/genproto/user_service/super_admins"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type superAdminRepo struct {
	db *pgxpool.Pool
}

func NewSuperAdminRepo(db *pgxpool.Pool) storage.SuperAdminRepo {
	return &superAdminRepo{
		db: db,
	}
}

func (s *superAdminRepo) Create(ctx context.Context, req *sp.CreateSuperAdmin) (*sp.SuperAdmin, error) {
	id := uuid.New()

	var externalId string

	row := s.db.QueryRow(ctx, `SELECT extra_id FROM super_admins ORDER BY created_at DESC LIMIT 1;`)

	err := row.Scan(&externalId)
	if err != nil {
		_, err = s.db.Exec(ctx, `
			INSERT INTO 
				super_admins(id, extra_id, fullname, phone, password, salary, months_worked, ielts_score)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, id, "SA00001", req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore)
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
		result := fmt.Sprintf("SA%s%d", zeros, number)
		_, err = s.db.Exec(ctx, `
			INSERT INTO 
				super_admins(id, extra_id, fullname, phone, password, salary, months_worked, ielts_score)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, id, result, req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore)

		if err != nil {
			return nil, err
		}
	}

	resp, err := s.GetById(ctx, &sp.SuperAdminPrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *superAdminRepo) Delete(ctx context.Context, req *sp.SuperAdminPrimaryKey) (*sp.Empty, error) {
	_, err := s.db.Exec(ctx, `DELETE FROM super_admins WHERE id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *superAdminRepo) GetAll(ctx context.Context, req *sp.GetListSuperAdminRequest) (*sp.GetListSuperAdminResponse, error) {
	resp := &sp.GetListSuperAdminResponse{}
	filter := ""

	if req.Search != "" {
		filter = ` AND fullname ILIKE '%` + req.Search + `%' `
	}
	
	rows, err := s.db.Query(ctx, `
	SELECT 
		id,
		extra_id,
		fullname,
		phone,
		password,
		salary,
		months_worked,
		ielts_score,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM
		super_admins
	WHERE
		TRUE ` + filter + `
	OFFSET
		$1
	LIMIT
		$2;`, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var superAdmin sp.SuperAdmin

		err := rows.Scan(
			&superAdmin.Id,
			&superAdmin.ExtraId,
			&superAdmin.FullName,
			&superAdmin.Phone,
			&superAdmin.Password,
			&superAdmin.Salary,
			&superAdmin.MonthsWorked,
			&superAdmin.IeltsScore,
			&superAdmin.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
	
		resp.SuperAdmins = append(resp.SuperAdmins, &superAdmin)
	}
	defer rows.Close()

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) FROM super_admins WHERE TRUE `+filter+`;`).Scan(&resp.Count)
	
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *superAdminRepo) GetById(ctx context.Context, req *sp.SuperAdminPrimaryKey) (*sp.SuperAdmin, error) {
	resp := &sp.SuperAdmin{}

	row := s.db.QueryRow(ctx, `
	SELECT 
		id,
		extra_id,
		fullname,
		phone,
		password,
		salary,
		months_worked,
		ielts_score,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM	
		super_admins
	WHERE
		id = $1;`, req.Id)

	if err := row.Scan(
		&resp.Id,
		&resp.ExtraId,
		&resp.FullName,
		&resp.Phone,
		&resp.Password,
		&resp.Salary,
		&resp.MonthsWorked,
		&resp.IeltsScore,
		&resp.CreatedAt,
	); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *superAdminRepo) Update(ctx context.Context, req *sp.UpdateSuperAdmin) (*sp.SuperAdmin, error) {
	_, err := s.db.Exec(ctx, `
	UPDATE
		super_admins	
	SET
		fullname = $2,
		phone = $3,
		password = $4,
		salary = $5,
		months_worked = $6,
		ielts_score = $7
	WHERE
		id = $1;`, req.Id, req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore)

	if err != nil {
		return nil, err
	}

	resp, err := s.GetById(ctx, &sp.SuperAdminPrimaryKey{Id: req.Id})

	if err != nil {
		return nil, err
	}
	return resp, nil
}
