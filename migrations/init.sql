CREATE
EXTENSION IF NOT EXISTS "uuid-ossp"; -- dong nay de tu dong sinh uuid

DROP TABLE IF EXISTS users;
create table users
(
    id    uuid DEFAULT uuid_generate_v4() primary key,
    phone varchar not null,
    email varchar null
);