CREATE TABLE to_dos (
    "id" int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "details" varchar NOT NULL,
    "status" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now())
)