CREATE TABLE IF NOT EXISTS "todos" (
                         "id" bigserial PRIMARY KEY,
                         "title" varchar NOT NULL,
                         "completed" boolean NOT NULL,
                         "item_order" integer NOT NULL,
                         "assigned_to" bigint,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "todos" ("id");

CREATE INDEX ON "todos" ("item_order");

CREATE INDEX ON "todos" ("assigned_to");