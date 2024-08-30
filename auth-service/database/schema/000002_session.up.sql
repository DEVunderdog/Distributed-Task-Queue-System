CREATE table "sessions" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "token" text NOT NULL,
    "refresh_token" text NOT NULL,
    "token_expires_at" timestamp with time zone NOT NULL,
    "refresh_token_expires_at" timestamp with time zone NOT NULL,
    "is_active" bool default true,
    "ip" varchar(100),
    "user_agent" varchar(255),
    "logged_out" timestamp with time zone,
    "created_at" timestamp with time zone not null default current_timestamp,
    "updated_at" timestamp with time zone
);

CREATE INDEX idx_user on "sessions" ("user_id");
CREATE INDEX idx_session_metadata on "sessions" ("ip", "user_agent");
