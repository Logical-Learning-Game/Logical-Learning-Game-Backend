CREATE TABLE IF NOT EXISTS map_configuration_item
(
    id                   BIGSERIAL PRIMARY KEY,
    map_configuration_id BIGINT  NOT NULL REFERENCES map_configuration (id),
    item_id              BIGINT  NOT NULL REFERENCES item (id),
    position_x           INTEGER NOT NULL,
    position_y           INTEGER NOT NULL
);