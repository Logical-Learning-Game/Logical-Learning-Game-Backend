CREATE TYPE map_direction AS ENUM
    (
        'up',
        'left',
        'down',
        'right'
        );

CREATE TABLE IF NOT EXISTS map
(
    id                     BIGSERIAL PRIMARY KEY,
    tile_array             INTEGER[][]   NOT NULL,
    height                 INTEGER       NOT NULL,
    width                  INTEGER       NOT NULL,
    start_player_direction map_direction NOT NULL
);