-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "users" (
    "id" uuid NOT NULL,
    "created_at" timestamptz DEFAULT now(),
    "deleted_at" timestamptz,
    "name" text,
    "email" text,
    PRIMARY KEY ("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "users" CASCADE;
