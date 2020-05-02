CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE user_details
(
    id   uuid DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    age  INT     NOT NULL
);
