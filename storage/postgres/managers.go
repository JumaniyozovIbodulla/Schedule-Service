package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	mn "user/genproto/user_service/managers"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type managerRepo struct {
	db *pgxpool.Pool
}

func NewManagerRepo(db *pgxpool.Pool) storage.ManagerRepo {
	return &managerRepo{
		db: db,
	}
}

func (m *managerRepo) Create(ctx context.Context, req *mn.CreateManager) (*mn.Manager, error) {
	resp := &mn.Manager{}
	id := uuid.New()

	var externalId string

	row := m.db.QueryRow(ctx, `SELECT extra_id FROM managers ORDER BY created_at DESC LIMIT 1;`)

	err := row.Scan(&externalId)
	if err != nil {
		_, err = m.db.Exec(ctx, `
			INSERT INTO 
				managers(id, extra_id, fullname, phone, password, salary, super_admin_id, branch_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, id, "M00001", req.FullName, req.Phone, req.Password, req.Salary, req.SuperAdminId, req.BranchId)
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
		result := fmt.Sprintf("M%s%d", zeros, number)
		_, err = m.db.Exec(ctx, `
			INSERT INTO 
				managers(id, extra_id, fullname, phone, password, salary, super_admin_id, branch_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, id, result, req.FullName, req.Phone, req.Password, req.Salary, req.SuperAdminId, req.BranchId)

		if err != nil {
			return nil, err
		}
	}

	resp, err = m.GetById(ctx, &mn.ManagerPrimaryKey{Id: id.String()})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *managerRepo) GetById(ctx context.Context, req *mn.ManagerPrimaryKey) (*mn.Manager, error) {
	resp := &mn.Manager{}

	row := m.db.QueryRow(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		salary, 
		super_admin_id, 
		branch_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM 
		managers 
	WHERE 
		id = $1;`, req.Id)

	err := row.Scan(&resp.Id, &resp.ExtraId, &resp.FullName, &resp.Phone, &resp.Password, &resp.Salary, &resp.SuperAdminId, &resp.BranchId, &resp.IsActive, &resp.CreatedAt)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *managerRepo) GetAll(ctx context.Context, req *mn.GetListManagersRequest) (*mn.GetListManagersResponse, error) {
	resp := &mn.GetListManagersResponse{}

	rows, err := m.db.Query(ctx, `
	SELECT 
		id, 
		extra_id, 
		fullname, 
		phone, 
		password,
		salary, 
		super_admin_id, 
		branch_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at 
	FROM 
		managers 
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
		var manager mn.Manager

		if err = rows.Scan(
			&manager.Id,
			&manager.ExtraId,
			&manager.FullName,
			&manager.Phone,
			&manager.Password,
			&manager.Salary,
			&manager.SuperAdminId,
			&manager.BranchId,
			&manager.IsActive,
			&manager.CreatedAt); err != nil {
			return nil, err
		}

		resp.Managers = append(resp.Managers, &manager)
	}

	err = m.db.QueryRow(ctx, `SELECT COUNT(*) FROM managers WHERE is_active = 0`).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *managerRepo) Update(ctx context.Context, req *mn.UpdateManager) (*mn.Manager, error) {
	resp := &mn.Manager{}
	_, err := m.db.Exec(ctx, `
	UPDATE 
		managers
	SET
		fullname = $2, 
		phone = $3, 
		password = $4, 
		salary = $5, 
		super_admin_id = $6,
		branch_id = $7
	WHERE
		id = $1;`, req.Id, req.FullName, req.Phone, req.Password, req.Salary, req.SuperAdminId, req.BranchId)

	if err != nil {
		return nil, err
	}
	resp, err = m.GetById(ctx, &mn.ManagerPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *managerRepo) Delete(ctx context.Context, req *mn.ManagerPrimaryKey) (*mn.Empty, error) {
	_, err := m.db.Exec(ctx, `
	UPDATE 
		managers 
	SET 
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT 
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return &mn.Empty{}, nil
}
