package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	gr "user/genproto/user_service/groups"
	"user/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

func mapLevelToPostgreSQL(level gr.CourseType) string {
	switch level {
	case gr.CourseType_beginner:
		return "beginner"
	case gr.CourseType_elementary:
		return "elementary"
	case gr.CourseType_pre_intermediate:
		return "pre-intermediate"
	case gr.CourseType_intermediate:
		return "intermediate"
	case gr.CourseType_upper_intermediate:
		return "upper-intermediate"
	case gr.CourseType_ielts_level_one:
		return "ielts-level-one"
	case gr.CourseType_ielts_level_two:
		return "ielts-level-two"
	default:
		return ""
	}
}

func mapPostgreSQLToCourseType(level string) gr.CourseType {
	switch level {
	case "beginner":
		return gr.CourseType_beginner
	case "elementary":
		return gr.CourseType_elementary
	case "pre-intermediate":
		return gr.CourseType_pre_intermediate
	case "intermediate":
		return gr.CourseType_intermediate
	case "upper-intermediate":
		return gr.CourseType_upper_intermediate
	case "ielts-level-one":
		return gr.CourseType_ielts_level_one
	case "ielts-level-two":
		return gr.CourseType_ielts_level_two
	default:
		return gr.CourseType(0)
	}
}

type groupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) storage.GroupRepo {
	return &groupRepo{
		db: db,
	}
}

func (g *groupRepo) Create(ctx context.Context, req *gr.CreateGroup) (*gr.Group, error) {
	id := uuid.New()

	level := mapLevelToPostgreSQL(req.Level)

	var externalId string

	row := g.db.QueryRow(ctx, `SELECT extra_id FROM groups ORDER BY created_at DESC LIMIT 1;`)

	err := row.Scan(&externalId)
	if err != nil {
		_, err = g.db.Exec(ctx, `
			INSERT INTO 
				groups(id, extra_id, name, level, months, number_of_students, branch_id)
			VALUES($1, $2, $3, $4, $5, $6, $7);`, id, "SA00001", req.Name, level, req.Months, req.NumberOfStudents, req.BranchId)
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
		_, err = g.db.Exec(ctx, `
			INSERT INTO 
				groups(id, extra_id, name, level, months, number_of_students, branch_id)
			VALUES($1, $2, $3, $4, $5, $6, $7);`, id, result, req.Name, level, req.Months, req.NumberOfStudents, req.BranchId)

		if err != nil {
			return nil, err
		}
	}

	group, err := g.GetById(ctx, &gr.GroupPrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}

	return group, nil
}

func (g *groupRepo) Delete(ctx context.Context, req *gr.GroupPrimaryKey) (*gr.Empty, error) {
	_, err := g.db.Exec(ctx, `
	UPDATE 
		groups 
	SET
		is_active = (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT
	WHERE 
		id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (g *groupRepo) GetAll(ctx context.Context, req *gr.GetListGroupsRequest) (*gr.GetListGroupsResponse, error) {
	resp := &gr.GetListGroupsResponse{}
	filter := ""

	if req.Search != "" {
		filter = ` AND name ILIKE '%` + req.Search + `%' `
	}

	rows, err := g.db.Query(ctx, ` 
	SELECT 
		id,
		extra_id,
		name, 
		level,
		months,
		number_of_students,
		branch_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM
		groups
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
			group gr.Group
			level pgtype.Text
		)

		if err = rows.Scan(
			&group.Id,
			&group.ExtraId,
			&group.Name,
			&level,
			&group.Months,
			&group.NumberOfStudents,
			&group.BranchId,
			&group.IsActive,
			&group.CreatedAt); err != nil {
			return nil, err
		}

		group.Level = mapPostgreSQLToCourseType(level.String)

		resp.Groups = append(resp.Groups, &group)
	}
	err = g.db.QueryRow(ctx, `SELECT COUNT(*) FROM groups WHERE TRUE `+filter+` AND is_active = 0;`).Scan(&resp.Count)

	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (g *groupRepo) GetById(ctx context.Context, req *gr.GroupPrimaryKey) (*gr.Group, error) {

	var (
		level pgtype.Text
		group gr.Group
	)

	row := g.db.QueryRow(ctx, `
	SELECT
		id,
		extra_id,
		name,
		level,
		months,
		number_of_students,
		branch_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM
		groups
	WHERE
		id = $1;`, req.Id)

	err := row.Scan(
		&group.Id,
		&group.ExtraId,
		&group.Name,
		&level,
		&group.Months,
		&group.NumberOfStudents,
		&group.BranchId,
		&group.IsActive,
		&group.CreatedAt)

	if err != nil {
		return nil, err
	}

	group.Level = mapPostgreSQLToCourseType(level.String)

	return &group, nil
}

func (g *groupRepo) Update(ctx context.Context, req *gr.UpdateGroup) (*gr.Group, error) {
	level := mapLevelToPostgreSQL(req.Level)

	_, err := g.db.Exec(ctx, `
	UPDATE 
		groups
	SET
		name = $2,
		level = $3,
		months = $4,
		number_of_students = $5,
		branch_id = $6
	WHERE
		id = $1;`,
		req.Id,
		req.Name,
		level,
		req.Months,
		req.NumberOfStudents,
		req.BranchId)

	if err != nil {
		return nil, err
	}

	resp, err := g.GetById(ctx, &gr.GroupPrimaryKey{Id: req.Id})

	if err != nil {
		return nil, err
	}
	return resp, nil
}
