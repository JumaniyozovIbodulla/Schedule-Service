package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	ad "user/genproto/user_service/administrators"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type adminRepo struct {
	db *pgxpool.Pool
}

func NewAdminRepo(db *pgxpool.Pool) storage.AdminRepo {
	return &adminRepo{
		db: db,
	}
}

func (a *adminRepo) Create(ctx context.Context, req *ad.CreateAdminstrator) (*ad.Adminstrator, error) {
	resp := &ad.Adminstrator{}
	id := uuid.New()

	var externalId string

	row := a.db.QueryRow(ctx, `SELECT extra_id FROM administrators ORDER BY created_at DESC LIMIT 1;`)

	err := row.Scan(&externalId)
	if err != nil {
		_, err = a.db.Exec(ctx, `
			INSERT INTO 
				administrators(id, extra_id, fullname, phone, password, salary, months_worked, ielts_score, branch_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`, id, "A00001", req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore, req.BranchId)
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
		result := fmt.Sprintf("A%s%d", zeros, number)
		_, err = a.db.Exec(ctx, `
			INSERT INTO 
				administrators(id, extra_id, fullname, phone, password, salary, months_worked, ielts_score, branch_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`, id, result, req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore, req.BranchId)

		if err != nil {
			return nil, err
		}
	}

	resp, err = a.GetById(ctx, &ad.AdminstratorPrimaryKey{Id: id.String()})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *adminRepo) GetById(ctx context.Context, req *ad.AdminstratorPrimaryKey) (*ad.Adminstrator, error) {
	resp := &ad.Adminstrator{}

	row := a.db.QueryRow(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		salary, 
		months_worked,
		ielts_score, 
		branch_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		administrators 
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
		&resp.BranchId,
		&resp.IsActive,
		&resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *adminRepo) GetAll(ctx context.Context, req *ad.GetListAdminstratorsRequest) (*ad.GetListAdminstratorsResponse, error) {
	resp := &ad.GetListAdminstratorsResponse{}

	rows, err := a.db.Query(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		salary, 
		months_worked,
		ielts_score, 
		branch_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		administrators 
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
		var admin ad.Adminstrator

		if err = rows.Scan(
			&admin.Id,
			&admin.ExtraId,
			&admin.FullName,
			&admin.Phone,
			&admin.Password,
			&admin.Salary,
			&admin.MonthsWorked,
			&admin.IeltsScore,
			&admin.BranchId,
			&admin.IsActive,
			&admin.CreatedAt); err != nil {
			return nil, err
		}

		resp.Adminstrators = append(resp.Adminstrators, &admin)
	}

	err = a.db.QueryRow(ctx, `SELECT COUNT(*) FROM administrators WHERE is_active = 0`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *adminRepo) Update(ctx context.Context, req *ad.UpdateAdminstrator) (*ad.Adminstrator, error) {
	resp := &ad.Adminstrator{}

	_, err := a.db.Exec(ctx, `
	UPDATE 
		administrators
	SET
		fullname = $2, 
		phone = $3, 
		password = $4, 
		salary = $5, 
		months_worked = $6,
		ielts_score = $7,
		branch_id = $8
	WHERE
		id = $1;`, req.Id, req.FullName, req.Phone, req.Password, req.Salary, req.MonthsWorked, req.IeltsScore, req.BranchId)

	if err != nil {
		return nil, err
	}
	resp, err = a.GetById(ctx, &ad.AdminstratorPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *adminRepo) Delete(ctx context.Context, req *ad.AdminstratorPrimaryKey) (*ad.Empty, error) {
	_, err := a.db.Exec(ctx, `
	UPDATE 
		administrators 
	SET 
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return &ad.Empty{}, nil
}
