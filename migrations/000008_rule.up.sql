CREATE TYPE rule_theme AS ENUM ('normal', 'conditional', 'loop');

CREATE TABLE IF NOT EXISTS rule
(
    name        VARCHAR(255) PRIMARY KEY,
    theme       rule_theme NOT NULL,
    description VARCHAR(255)
);