-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- dong nay de tu dong sinh uuid
DROP TABLE IF EXISTS users;

create table public.users (
    id uuid DEFAULT uuid_generate_v4() primary key,
    phone varchar not null,
    email varchar null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table public.rooms(
    id bigserial primary key,
    name varchar not null,
    "seats" varchar [] not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table public.movies(
    id uuid DEFAULT uuid_generate_v4() primary key,
    title varchar not null,
    content text null,
    videos jsonb null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd