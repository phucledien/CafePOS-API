-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "orders" (
  "id" uuid NOT NULL,
  "table_id" uuid references tables,
  "total_payment" int,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  PRIMARY KEY ("id")
) with (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "orders" CASCADE;
