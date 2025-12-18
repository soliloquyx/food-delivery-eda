package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/soliloquyx/food-delivery-eda/internal/order/app"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) orderRepo {
	return orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(
	ctx context.Context,
	orderID uuid.UUID,
	in app.PlaceOrderInput,
) (app.PlaceOrderResult, error) {
	return app.PlaceOrderResult{
		OrderID: orderID,
		Status:  app.StatusConfirmed,
	}, nil
}
