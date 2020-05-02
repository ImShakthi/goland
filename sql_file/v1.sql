CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_details
(
    id   uuid DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    age  INT     NOT NULL
);
