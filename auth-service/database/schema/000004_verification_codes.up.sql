CREATE TABLE "verification_codes" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "code" varchar(50) NOT NULL,
    "expires_at" timestamp with time zone NOT NULL,
    "is_used" bool NOT NULL default false,
    "created_at" timestamp with time zone NOT NULL default current_timestamp,
    "updated_at" timestamp with time zone
);

ALTER TABLE "verification_codes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
CREATE INDEX idx_user_is_used on "verification_codes" ("user_id", "is_used");