-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- dong nay de tu dong sinh uuid
DROP TABLE IF EXISTS public.constants;

create table public.constants (
    "code" varchar primary key,
    "value" text not null
);

DROP TABLE IF EXISTS public.users;

create table public.users (
    id uuid DEFAULT uuid_generate_v4() primary key,
    "loyalty_id" int NULL,
    "email" varchar NULL,
    "phone" varchar NOT NULL,
    "cur_original_id" varchar NULL,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS public.rooms;

create table public.rooms(
    id bigserial primary key,
    "name" varchar not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS public.seats;

create table public.seats(
    id uuid DEFAULT uuid_generate_v4() primary key,
    seat_code varchar not null,
    room_id integer not null,
    seat_type varchar not null,
    seat_order integer not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS public.movies;

create table public.movies(
    id uuid DEFAULT uuid_generate_v4() primary key,
    title varchar not null,
    content text null,
    videos jsonb null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS public.slots;

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

DROP TABLE IF EXISTS public.orders;

create table public.orders (
    id varchar(30) not null primary key,
    user_id uuid not null,
    slot_id uuid not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS public.slot_seats;

create table public.slot_seats(
    id uuid DEFAULT uuid_generate_v4() primary key,
    seat_id uuid not null,
    slot_id uuid not null,
    order_id varchar null,
    total_pay float null,
    "status" varchar not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX slot_seats_slot_id_seat_id_u_idx ON public.slot_seats(slot_id, seat_id);

CREATE INDEX orders_user_id_idx ON public.orders (user_id);

CREATE INDEX orders_slot_id_idx ON public.orders (slot_id);

CREATE UNIQUE INDEX seats_seat_code_room_id_u_idx ON public.seats (seat_code, room_id);

CREATE UNIQUE INDEX users_loyalty_id_u_idx ON public.users(loyalty_id);

CREATE UNIQUE INDEX users_phone_u_idx ON public.users(phone);

CREATE UNIQUE INDEX users_email_u_idx ON public.users(email);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd