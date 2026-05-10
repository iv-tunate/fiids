-- +goose Up
ALTER TABLE feeds
ALTER COLUMN updated_at DROP DEFAULT;

-- +goose Down
ALTER TABLE feeds
ALTER COLUMN updated_at SET DEFAULT now();