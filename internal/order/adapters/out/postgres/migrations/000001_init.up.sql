BEGIN;

CREATE TABLE IF NOT EXISTS orders (
    id uuid PRIMARY KEY DEFAULT uuidv7(),
    user_id uuid NOT NULL,
    restaurant_id uuid NOT NULL,
    fulfillment_type text NOT NULL,
    delivery_address text,
    delivery_comment text NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    status text NOT NULL DEFAULT 'pending',
    CONSTRAINT order_valid_fulfillment_type
        CHECK (
            (fulfillment_type = 'pickup' AND delivery_address IS NULL)
            OR
            (fulfillment_type = 'delivery' AND delivery_address IS NOT NULL)
        ),
    CONSTRAINT order_valid_status CHECK (status IN ('pending', 'confirmed', 'cancelled'))
);

CREATE TABLE IF NOT EXISTS order_items (
    id uuid PRIMARY KEY DEFAULT uuidv7(),
    item_id uuid NOT NULL,
    order_id uuid NOT NULL REFERENCES orders(id),
    comment text NOT NULL DEFAULT '',
    quantity integer NOT NULL,
    CONSTRAINT positive_order_item_quantity CHECK (quantity > 0),
    UNIQUE (order_id, item_id)
);

COMMIT;
