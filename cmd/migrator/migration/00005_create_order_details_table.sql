-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "order_details" (
  "id" uuid NOT NULL,
  "order_id" uuid references orders,
  "drink_id" uuid references drinks,
  "quantity" int,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  PRIMARY KEY ("id")
) with (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "order_details" CASCADE;
