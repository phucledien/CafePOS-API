-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "drinks" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" text,
  "price" int,
  PRIMARY KEY ("id")
) with (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "drinks" CASCADE;
