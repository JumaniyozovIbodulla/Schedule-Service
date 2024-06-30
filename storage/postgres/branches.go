package postgres

import (
	"context"
	br "user/genproto/user_service/branches"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepo {
	return &branchRepo{
		db: db,
	}
}

func (b *branchRepo) Create(ctx context.Context, req *br.CreateBranch) (*br.Branch, error) {
	id := uuid.New()

	_, err := b.db.Exec(ctx, `
		INSERT INTO
			branches(id, name, location, super_admin_id)
		VALUES(
		$1, 
		$2, 
		ST_SetSRID(ST_MakePoint($3, $4), 4326),
		$5);`,
		id,
		req.Name,
		req.Location.Latitude,
		req.Location.Longitude,
		req.SuperAdminId)

	if err != nil {
		return nil, err
	}

	branch, err := b.GetById(ctx, &br.BranchePrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}

	return branch, nil
}

func (b *branchRepo) Delete(ctx context.Context, req *br.BranchePrimaryKey) (*br.Empty, error) {
	_, err := b.db.Exec(ctx, `
	UPDATE 
		branches 
	SET
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (b *branchRepo) GetAll(ctx context.Context, req *br.GetListBranchesRequest) (*br.GetListBranchesResponse, error) {
	resp := &br.GetListBranchesResponse{}
	filter := ""

	if req.Search != "" {
		filter = ` AND name ILIKE '%` + req.Search + `%' `
	}

	rows, err := b.db.Query(ctx, ` 
	SELECT 
		id,
		name,
		ST_X(location) AS latitude, 
		ST_Y(location) AS longitude,
		super_admin_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM
		branches
	WHERE TRUE `+filter+` AND is_active = 0
	OFFSET
		$1
	LIMIT
		$2;`, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			branch              br.Branch
			longitude, latitude float64
		)

		if err = rows.Scan(
			&branch.Id,
			&branch.Name,
			&latitude,
			&longitude,
			&branch.SuperAdminId,
			&branch.IsActive,
			&branch.CreatedAt); err != nil {
			return nil, err
		}

		branch.Location = &br.Location{
			Latitude:  latitude,
			Longitude: longitude,
		}

		resp.Branches = append(resp.Branches, &branch)
	}
	err = b.db.QueryRow(ctx, `SELECT COUNT(*) FROM branches WHERE TRUE `+filter+` AND is_active = 0;`).Scan(&resp.Count)

	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (b *branchRepo) GetById(ctx context.Context, req *br.BranchePrimaryKey) (*br.Branch, error) {

	branch := &br.Branch{}

	row := b.db.QueryRow(ctx, `
	SELECT
		id,
		name,
		ST_X(location) AS latitude, 
		ST_Y(location) AS longitude,
		super_admin_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM
		branches
	WHERE
		id = $1;`, req.Id)

	var lat, long float64
	err := row.Scan(
		&branch.Id,
		&branch.Name,
		&lat,
		&long,
		&branch.SuperAdminId,
		&branch.IsActive,
		&branch.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	branch.Location = &br.Location{
		Latitude:  lat,
		Longitude: long,
	}

	return branch, nil
}

func (b *branchRepo) Update(ctx context.Context, req *br.UpdateBranch) (*br.Branch, error) {

	_, err := b.db.Exec(ctx, `
	UPDATE 
		branches
	SET
		name = $2,
		location = ST_SetSRID(ST_MakePoint($3, $4), 4326),
		super_admin_id = $5
	WHERE
		id = $1;`,
		req.Id,
		req.Name,
		req.Location.Latitude,
		req.Location.Longitude,
		req.SuperAdminId)

	if err != nil {
		return nil, err
	}

	resp, err := b.GetById(ctx, &br.BranchePrimaryKey{Id: req.Id})

	if err != nil {
		return nil, err
	}
	return resp, nil
}
