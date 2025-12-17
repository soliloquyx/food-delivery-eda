FROM migrate/migrate:v4.19.1
COPY internal/order/adapters/out/postgres/migrations migrations/
