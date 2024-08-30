CREATE table "jwtkeys" (
    "id" bigserial PRIMARY KEY,
    "public_key" text NOT NULL,
    "private_key" text NOT NULL,
    "algorithm" text NOT NULL,
    "is_active" bool default true,
    "expires_at" timestamp with time zone,
    "created_at" timestamp with time zone not null default current_timestamp,
    "updated_at" timestamp with time zone
);

CREATE INDEX idx_active_jwtkeys ON "jwtkeys" ("is_active");
