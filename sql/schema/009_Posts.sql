-- +goose Up

CREATE TABLE posts(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    published_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    url TEXT NOT NULL UNIQUE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);