package domain

import (
	"context"
	"fmt"
	"math"

	"github.com/Beelzebub0/go-crud-boilerplate/src/business/entity"
	"github.com/pkg/errors"
)

func (dom *domain) sqlCreateArea(ctx context.Context, params entity.AreaInput) (entity.Area, error) {
	result := entity.Area{}

	db, err := dom.sql.Connect()
	if err != nil {
		return result, err
	}

	defer db.Close()

	res, err := db.ExecContext(ctx, _sqlCreateArea,
		params.Name,
		// params.Address,
		params.Latitude,
		params.Longitude,
		params.CodeArea,
		params.Notes,
	)
	if err != nil {
		return result, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return result, err
	}

	result, err = dom.sqlGetAreaByID(ctx, lid)
	if err != nil {
		return result, err
	}

	return result, err
}

func (dom *domain) sqlGetAreaByID(ctx context.Context, aid int64) (entity.Area, error) {
	result := entity.Area{}

	db, err := dom.sql.Connect()
	if err != nil {
		return result, err
	}

	defer db.Close()

	err = db.QueryRowContext(ctx, _sqlGetAreaByID, aid).Scan(
		&result.ID,
		&result.Name,
		// &result.Address,
		&result.Latitude,
		&result.Longitude,
		&result.CodeArea,
		&result.Notes,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return result, err
	}

	return result, err
}

func (dom *domain) sqlGetArea(ctx context.Context, params entity.AreaParams) ([]entity.Area, entity.Pagination, error) {
	result := []entity.Area{}
	pagination := entity.Pagination{}

	db, err := dom.sql.Connect()
	if err != nil {
		return result, pagination, err
	}

	defer db.Close()

	query := params.CreateMySQLQuery(_sqlGetArea, true)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return result, pagination, err
	}

	defer rows.Close()

	for rows.Next() {
		var r entity.Area
		err = rows.Scan(
			&r.ID,
			&r.Name,
			// &r.Address,
			&r.Latitude,
			&r.Longitude,
			&r.CodeArea,
			&r.Notes,
			&r.CreatedAt,
			&r.UpdatedAt,
		)
		if err != nil {
			fmt.Print(err)
			return result, pagination, err
		}
		result = append(result, r)
	}

	countquery := params.CreateMySQLQuery(_sqlGetAreaCount, false)

	err = db.QueryRowContext(ctx, countquery).Scan(&pagination.TotalElements)
	if err != nil {
		return result, pagination, errors.Wrap(err, "Error Executing SQL Query")
	}

	pagination.CurrentElements = int64(len(result))
	pagination.CurrentPage = params.Page
	pagination.TotalPages = int64(math.Ceil(float64(pagination.TotalElements) / float64(params.Limit)))

	return result, pagination, err
}

func (dom *domain) sqlUpdateArea(ctx context.Context, params entity.UpdateAreaInput) (entity.Area, error) {
	result := entity.Area{}

	currVal, err := dom.sqlGetAreaByID(ctx, params.ID)
	if err != nil {
		return result, err
	}

	params.SetDefaultValue(currVal)

	db, err := dom.sql.Connect()
	if err != nil {
		return result, err
	}

	_, err = db.ExecContext(ctx, _sqlUpdateArea,
		params.Name,
		// params.Address,
		params.CodeArea,
		params.Notes,
		params.Latitude,
		params.Longitude,
		params.ID,
	)
	if err != nil {
		return result, err
	}

	result, err = dom.sqlGetAreaByID(ctx, params.ID)
	if err != nil {
		return result, err
	}

	return result, err
}

func (dom *domain) sqlDeleteArea(ctx context.Context, aid int64) error {
	db, err := dom.sql.Connect()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, _sqlDeleteArea, aid)
	if err != nil {
		return err
	}

	return err
}
