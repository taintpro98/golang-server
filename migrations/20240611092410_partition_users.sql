-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

create table public.users (
  id uuid DEFAULT uuid_generate_v4() primary key,
  "loyalty_id" int NULL,
  "email" varchar NULL,
  "phone" varchar NOT NULL,
  "cur_original_id" varchar NULL,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
) partition by range(created_at);

CREATE UNIQUE INDEX IF NOT EXISTS users_phone_u_idx ON public.users USING HASH (phone);

CREATE UNIQUE INDEX IF NOT EXISTS users_loyalty_id_u_idx ON public.users USING HASH (loyalty_id);

-- Creating a function to generate partition tables dynamically
CREATE
OR REPLACE FUNCTION create_users_partition(start_date DATE, end_date DATE) RETURNS VOID AS $ $ DECLARE curr_date DATE := start_date;

BEGIN WHILE curr_date < end_date LOOP EXECUTE format(
  'CREATE TABLE IF NOT EXISTS public."users_%s" PARTITION OF public.users FOR VALUES FROM (%L) TO (%L)',
  to_char(curr_date, 'YYYYMMMM'),
  curr_date,
  curr_date + INTERVAL '1 month'
);

END LOOP;

END;

$ $ LANGUAGE plpgsql;

SELECT
  create_users_partition('2024-03-01' :: DATE, '2025-01-01' :: DATE);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

DROP TABLE IF EXISTS users;

-- +goose StatementEnd