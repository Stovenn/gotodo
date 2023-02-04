CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "full_name" varchar NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "users" ("id");

CREATE INDEX ON "todos" ("id");

CREATE INDEX ON "todos" ("item_order");

ALTER TABLE "todos" ADD FOREIGN KEY ("assigned_to") REFERENCES "users" ("id");