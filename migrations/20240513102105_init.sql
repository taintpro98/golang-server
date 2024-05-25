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
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table public.rooms(
    id bigserial primary key,
    name varchar not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table public.seats(
    id uuid DEFAULT uuid_generate_v4() primary key,
    seat_code varchar not null,
    room_id integer not null,
    seat_type varchar not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table public.movies(
    id uuid DEFAULT uuid_generate_v4() primary key,
    title varchar not null,
    content text null,
    videos jsonb null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table public.slots(
    id uuid DEFAULT uuid_generate_v4() primary key,
    room_id integer not null,
    movie_id uuid not null,
    start_time timestamp not null,
    end_time timestamp null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table public.orders (
    id varchar(30) not null primary key,
    user_id uuid not null,
    slot_id uuid not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

create table public.order_seats (
    id uuid DEFAULT uuid_generate_v4() primary key,
    order_id varchar(30) not null,
    seat_id uuid not null,
    total_pay float not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX orders_user_id_idx ON public.orders (user_id);

CREATE INDEX orders_slot_id_idx ON public.orders (slot_id);

CREATE UNIQUE INDEX seats_seat_code_room_id_u_idx ON public.seats (seat_code, room_id);

CREATE UNIQUE INDEX order_seats_order_id_seat_id_u_idx ON public.order_seats (order_id, seat_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd