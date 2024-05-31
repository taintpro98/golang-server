-- +goose Up
-- +goose StatementBegin
insert into
  public.constants (code, "value")
values
  ('users_num', '0');

INSERT INTO
  public.movies (
    id,
    title,
    "content",
    videos,
    created_at,
    updated_at,
    deleted_at
  )
VALUES
  (
    '48ab1d8a-1273-42f8-b46e-7bb16e995794',
    'Con buom xinh',
    'Mot nguoi di xuyen qua khong thoi gian cung con buom xinh',
    NULL,
    '2024-05-25 09:03:56.178',
    '2024-05-25 09:03:56.178',
    '2024-05-25 17:03:56.178'
  ),
  (
    'a021fd03-2600-493c-8344-b469d24a6e3e',
    'Nguoi bat tu',
    'Mot nguoi di xuyen qua khong',
    NULL,
    '2024-05-25 09:27:21.576',
    '2024-05-25 09:27:21.576',
    '2024-05-25 17:27:21.575'
  );

INSERT INTO
  public.rooms ("name", created_at, updated_at, deleted_at)
VALUES
  (
    'Room 1',
    '2024-05-25 09:26:30.524',
    '2024-05-25 09:26:30.524',
    '2024-05-25 17:26:30.523'
  ),
  (
    'Room 2',
    '2024-05-25 09:26:48.159',
    '2024-05-25 09:26:48.159',
    '2024-05-25 17:26:48.158'
  );

INSERT INTO
  public.slots (
    id,
    room_id,
    movie_id,
    start_time,
    end_time,
    created_at,
    updated_at,
    deleted_at
  )
VALUES
  (
    '527e5d2c-3703-4bba-8ea3-0d465bb945b7',
    1,
    '48ab1d8a-1273-42f8-b46e-7bb16e995794',
    '2024-05-25 09:04:21.267',
    '1970-01-01 00:00:00.000',
    '2024-05-25 09:04:21.267',
    '2024-05-25 09:04:21.267',
    '2024-05-25 17:04:21.266'
  ),
  (
    '4bdacdba-16b5-4160-bec2-4df443611f13',
    1,
    '48ab1d8a-1273-42f8-b46e-7bb16e995794',
    '2024-12-15 03:52:03.000',
    NULL,
    '2024-05-25 09:08:37.450',
    '2024-05-25 09:08:37.450',
    '2024-05-25 17:08:37.450'
  ),
  (
    '17e0f305-c7ca-446d-9b12-689a4dd6824f',
    1,
    '48ab1d8a-1273-42f8-b46e-7bb16e995794',
    '2024-12-15 03:52:03.000',
    NULL,
    '2024-05-25 09:26:52.715',
    '2024-05-25 09:26:52.715',
    '2024-05-25 17:26:52.715'
  ),
  (
    'fe0e7e9e-e3c6-402a-9f5e-ebd0f077b12e',
    2,
    '48ab1d8a-1273-42f8-b46e-7bb16e995794',
    '2024-12-15 03:52:03.000',
    NULL,
    '2024-05-25 09:27:01.657',
    '2024-05-25 09:27:01.657',
    '2024-05-25 17:27:01.656'
  ),
  (
    '959df217-c984-4aef-b3f8-fbc509fc1e64',
    2,
    'a021fd03-2600-493c-8344-b469d24a6e3e',
    '2024-12-15 03:52:03.000',
    NULL,
    '2024-05-25 09:27:31.667',
    '2024-05-25 09:27:31.667',
    '2024-05-25 17:27:31.666'
  );

insert into
  public.seats (room_id, seat_code, seat_type, seat_order)
values
  (1, 'A1', 'normal', 1),
  (1, 'A2', 'normal', 2),
  (1, 'A3', 'normal', 3),
  (1, 'B1', 'vip', 4),
  (1, 'B2', 'vip', 5),
  (1, 'B3', 'vip', 6);

insert into
  public.users (id, phone, email)
values
  (
    '566b2ae2-5837-4b20-a030-b6825308c288',
    '012345678',
    'alice@gmail.com'
  ),
  (
    '73122f85-50db-4f1d-a618-672f0609fe61',
    '023456781',
    'bob@gmail.com'
  ),
  (
    '719565e7-4afc-4f21-9fbf-836053662dde',
    '034567812',
    'clover@gmail.com'
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd