package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mirhijinam/wb-l0/internal/models"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) Repo {
	return Repo{
		pool: pool,
	}
}

func (r *Repo) Save(ctx context.Context, order models.OrderData) error {
	const op = `repo.Save`

	stmt := `
		INSERT INTO orders (order_uid, data)
		VALUES ($1, $2)
	`

	_, err := r.pool.Exec(ctx, stmt, order.OrderUid, order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) PutIntoCache(ctx context.Context) (map[string]models.OrderData, error) {
	const op = `repo.PutIntoCache`

	stmt := `
		SELECT * FROM orders
	`

	rows, err := r.pool.Query(ctx, stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	m := make(map[string]models.OrderData, 256)
	for rows.Next() {
		var data models.OrderData

		if err = rows.Scan(&data); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		m[data.OrderUid] = data
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return m, nil
}
