-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

CREATE TABLE IF NOT EXISTS public."hail_driver_location" (
  "id" bigserial,
  "driver_id" uuid NOT NULL,
  "order_id" TEXT NOT NULL,
  "lat" FLOAT NOT NULL,
  "lng" FLOAT NOT NULL,
  "time" timestamp NOT NULL,
  "time_delta" int NOT NULL,
  "speed" FLOAT,
  "speed_acc" FLOAT,
  "bearing" FLOAT,
  "bearing_acc" FLOAT,
  "horizontal_acc" FLOAT,
  "vertical_acc" FLOAT,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
) PARTITION BY RANGE(time);

CREATE INDEX IF NOT EXISTS hail_driver_location_idx ON public.hail_driver_location (driver_id, order_id, time);

-- Creating a function to generate partition tables dynamically
CREATE
OR REPLACE FUNCTION create_partitions(start_date DATE, end_date DATE) RETURNS VOID AS $ $ DECLARE curr_date DATE := start_date;

BEGIN WHILE curr_date < end_date LOOP EXECUTE format(
  '
            CREATE TABLE IF NOT EXISTS public."hail_driver_location_%s" PARTITION OF public."hail_driver_location"
            FOR VALUES FROM (%L) TO (%L);',
  to_char(curr_date, 'YYYYMMDD'),
  curr_date,
  curr_date + INTERVAL '1 day'
);

curr_date := curr_date + INTERVAL '1 day';

END LOOP;

END;

$ $ LANGUAGE plpgsql;

SELECT
  create_partitions('2024-03-01' :: DATE, '2025-01-01' :: DATE);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

DROP TABLE IF EXISTS hail_driver_location;

-- +goose StatementEnd