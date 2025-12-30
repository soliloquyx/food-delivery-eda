package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/soliloquyx/food-delivery-eda/internal/order/order"
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
	in order.PlaceOrderInput,
) (order.PlaceOrderResult, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return order.PlaceOrderResult{}, err
	}
	defer tx.Rollback(ctx)

	var (
		status          order.Status
		createdAt       time.Time
		deliveryAddr    *string
		deliveryComment string
	)

	if in.Delivery != nil {
		deliveryAddr = &in.Delivery.Address
		deliveryComment = in.Delivery.Comment
	}

	if err := tx.QueryRow(
		ctx, `
		INSERT INTO orders(
			id,
			user_id,
			restaurant_id,
			fulfillment_type,
			delivery_address,
			delivery_comment
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING status, created_at;
		`, orderID, in.UserID, in.RestaurantID, in.FulfillmentType, deliveryAddr, deliveryComment,
	).Scan(&status, &createdAt); err != nil {
		fmt.Printf("%+v\n", in)
		return order.PlaceOrderResult{}, err
	}

	if _, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"order_items"},
		[]string{"item_id", "order_id", "quantity", "comment"},
		pgx.CopyFromSlice(len(in.Items), func(i int) ([]any, error) {
			return []any{in.Items[i].ItemID, orderID, in.Items[i].Quantity, in.Items[i].Comment}, nil
		}),
	); err != nil {
		return order.PlaceOrderResult{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return order.PlaceOrderResult{}, err
	}

	return order.PlaceOrderResult{
		OrderID:   orderID,
		Status:    status,
		CreatedAt: createdAt,
	}, nil
}
