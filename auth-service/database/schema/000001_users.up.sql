CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "email" varchar(100) NOT NULL,
    "hashed_password" text NOT NULL,
    "email_verified" bool NOT NULL,
    "created_at" timestamp with time zone not null default current_timestamp,
    "updated_at" timestamp with time zone
);

CREATE INDEX idx_email ON "users" ("email");