-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- dong nay de tu dong sinh uuid
DROP TABLE IF EXISTS users;

create table users (
    id uuid DEFAULT uuid_generate_v4() primary key,
    phone varchar not null,
    email varchar null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table posts(
    id uuid DEFAULT uuid_generate_v4() primary key,
    user_id uuid not null,
    title varchar not null,
    content varchar null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd