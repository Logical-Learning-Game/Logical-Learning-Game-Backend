CREATE TABLE IF NOT EXISTS map_configuration_for_player
(
    id                   BIGSERIAL PRIMARY KEY,
    map_configuration_id BIGINT       NOT NULL REFERENCES map_configuration (id),
    player_id            VARCHAR(255) NOT NULL REFERENCES player (player_id),
    is_pass              BOOLEAN      NOT NULL
);