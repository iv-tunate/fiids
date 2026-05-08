-- +goose Up

ALTER TABLE users
    ADD CONSTRAINT users_email_unique UNIQUE(email);

ALTER TABLE users
    ALTER COLUMN created_at TYPE TIMESTAMPTZ
    USING created_at AT TIME ZONE 'UTC';

ALTER TABLE users
    ALTER COLUMN updated_at TYPE TIMESTAMPTZ
    USING updated_at AT TIME ZONE 'UTC';

ALTER TABLE users
    ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE users
    ALTER COLUMN updated_at SET DEFAULT now();

-- +goose Down

ALTER TABLE users
    DROP CONSTRAINT users_email_unique;