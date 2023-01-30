CREATE TABLE IF NOT EXISTS map_configuration_door
(
    id                   BIGSERIAL PRIMARY KEY,
    map_configuration_id BIGINT        NOT NULL REFERENCES map_configuration (id),
    door_id              BIGINT        NOT NULL REFERENCES door (id),
    position_x           INTEGER       NOT NULL,
    position_y           INTEGER       NOT NULL,
    door_direction       map_direction NOT NULL
);