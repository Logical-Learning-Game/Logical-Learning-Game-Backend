CREATE TYPE command_node_type AS ENUM
    (
        'conditional_a',
        'conditional_b',
        'conditional_c',
        'conditional_d',
        'conditional_e',
        'up',
        'left',
        'down',
        'right'
        );

CREATE TABLE IF NOT EXISTS command_node
(
    id                 BIGSERIAL PRIMARY KEY,
    play_history_id    BIGINT            NOT NULL REFERENCES play_history (id),
    type               command_node_type NOT NULL,
    in_game_position_x REAL              NOT NULL,
    in_game_position_y REAL              NOT NULL
);