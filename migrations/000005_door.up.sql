CREATE TYPE door_type AS ENUM ('door_no_key', 'door_a', 'door_b', 'door_c');

CREATE TABLE IF NOT EXISTS door
(
    id     BIGSERIAL PRIMARY KEY,
    name   VARCHAR(255) NOT NULL,
    active BOOLEAN      NOT NULL,
    type   door_type    NOT NULL
);