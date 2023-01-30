CREATE TABLE IF NOT EXISTS game_session
(
    id                   BIGSERIAL PRIMARY KEY,
    player_id            VARCHAR(255) NOT NULL REFERENCES player (player_id),
    map_configuration_id BIGINT       NOT NULL REFERENCES map_configuration (id),
    start_datetime       TIMESTAMPTZ  NOT NULL DEFAULT (now()),
    end_datetime         TIMESTAMPTZ
);