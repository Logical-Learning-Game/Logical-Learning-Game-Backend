CREATE TYPE item_type AS ENUM ('key_a', 'key_b', 'key_c');

CREATE TABLE IF NOT EXISTS item
(
    id     BIGSERIAL PRIMARY KEY,
    name   VARCHAR(255) NOT NULL,
    active BOOLEAN      NOT NULL,
    type   item_type    NOT NULL
);