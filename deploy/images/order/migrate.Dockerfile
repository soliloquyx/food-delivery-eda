FROM migrate/migrate:v4.19.1

COPY internal/order/adapters/postgres/migrations migrations/
