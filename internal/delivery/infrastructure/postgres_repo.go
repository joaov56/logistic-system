package infrastructure

import (
	"context"
	"database/sql"
	"logistic-system/internal/delivery/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, delivery *domain.Delivery) error {
	query := `
		INSERT INTO deliveries (id, order_id, customer_id, address, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		delivery.ID,
		delivery.OrderID,
		delivery.CustomerID,
		delivery.Address,
		delivery.Status,
		delivery.CreatedAt,
		delivery.UpdatedAt,
	)
	return err
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*domain.Delivery, error) {
	query := `
		SELECT id, order_id, customer_id, address, status, created_at, updated_at, delivered_at
		FROM deliveries
		WHERE id = $1
	`
	var delivery domain.Delivery
	var deliveredAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&delivery.ID,
		&delivery.OrderID,
		&delivery.CustomerID,
		&delivery.Address,
		&delivery.Status,
		&delivery.CreatedAt,
		&delivery.UpdatedAt,
		&deliveredAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewDomainError("delivery not found")
		}
		return nil, err
	}

	if deliveredAt.Valid {
		delivery.DeliveredAt = &deliveredAt.Time
	}

	return &delivery, nil
}

func (r *PostgresRepository) Update(ctx context.Context, delivery *domain.Delivery) error {
	query := `
		UPDATE deliveries
		SET status = $1, updated_at = $2, delivered_at = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query,
		delivery.Status,
		delivery.UpdatedAt,
		delivery.DeliveredAt,
		delivery.ID,
	)
	return err
}

func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM deliveries WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *PostgresRepository) List(ctx context.Context, filter map[string]interface{}) ([]*domain.Delivery, error) {
	query := `
		SELECT id, order_id, customer_id, address, status, created_at, updated_at, delivered_at
		FROM deliveries
	`
	args := []interface{}{}
	whereClause := ""

	if len(filter) > 0 {
		whereClause = "WHERE "
		i := 1
		for key, value := range filter {
			if i > 1 {
				whereClause += " AND "
			}
			whereClause += key + " = $" + string(rune('0'+i))
			args = append(args, value)
			i++
		}
	}

	rows, err := r.db.QueryContext(ctx, query+whereClause, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deliveries []*domain.Delivery
	for rows.Next() {
		var delivery domain.Delivery
		var deliveredAt sql.NullTime

		err := rows.Scan(
			&delivery.ID,
			&delivery.OrderID,
			&delivery.CustomerID,
			&delivery.Address,
			&delivery.Status,
			&delivery.CreatedAt,
			&delivery.UpdatedAt,
			&deliveredAt,
		)
		if err != nil {
			return nil, err
		}

		if deliveredAt.Valid {
			delivery.DeliveredAt = &deliveredAt.Time
		}

		deliveries = append(deliveries, &delivery)
	}

	return deliveries, nil
} 