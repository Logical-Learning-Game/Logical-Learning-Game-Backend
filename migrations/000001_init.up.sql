CREATE TABLE IF NOT EXISTS player
(
    player_id VARCHAR(255) PRIMARY KEY,
    email     VARCHAR(255) DEFAULT NULL,
    name      VARCHAR(255) DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS login_log
(
    player_id VARCHAR(255) NOT NULL REFERENCES player (player_id),
    login_at  TIMESTAMPTZ  NOT NULL DEFAULT (now())
);

CREATE TYPE map_direction AS ENUM
    (
        'up',
        'left',
        'down',
        'right'
        );

CREATE TABLE IF NOT EXISTS map
(
    id                      BIGSERIAL PRIMARY KEY,
    tile_array              INTEGER[]     NOT NULL,
    height                  INTEGER       NOT NULL,
    width                   INTEGER       NOT NULL,
    start_player_direction  map_direction NOT NULL,
    start_player_position_x INTEGER       NOT NULL,
    start_player_position_y INTEGER       NOT NULL,
    goal_position_x         INTEGER       NOT NULL,
    goal_position_y         INTEGER       NOT NULL
);

CREATE TYPE item_type AS ENUM ('key_a', 'key_b', 'key_c');

CREATE TABLE IF NOT EXISTS item
(
    id     BIGSERIAL PRIMARY KEY,
    name   VARCHAR(255) NOT NULL,
    active BOOLEAN      NOT NULL,
    type   item_type    NOT NULL
);

CREATE TYPE door_type AS ENUM ('door_no_key', 'door_a', 'door_b', 'door_c');

CREATE TABLE IF NOT EXISTS door
(
    id     BIGSERIAL PRIMARY KEY,
    name   VARCHAR(255) NOT NULL,
    active BOOLEAN      NOT NULL,
    type   door_type    NOT NULL
);

CREATE TABLE IF NOT EXISTS world
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TYPE map_difficulty AS ENUM ('easy', 'medium', 'hard');

CREATE TABLE IF NOT EXISTS map_configuration
(
    id                            BIGSERIAL PRIMARY KEY,
    world_id                      BIGINT         NOT NULL REFERENCES world (id),
    map_id                        BIGINT         NOT NULL REFERENCES map (id),
    config_name                   VARCHAR(255)   NOT NULL,
    map_image_path                VARCHAR(255),
    difficulty                    map_difficulty NOT NULL,
    star_requirement              INTEGER        NOT NULL,
    least_solvable_command_gold   INTEGER        NOT NULL,
    least_solvable_command_silver INTEGER        NOT NULL,
    least_solvable_command_bronze INTEGER        NOT NULL,
    least_solvable_action_gold    INTEGER        NOT NULL,
    least_solvable_action_silver  INTEGER        NOT NULL,
    least_solvable_action_bronze  INTEGER        NOT NULL,
    created_at                    TIMESTAMPTZ    NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS rule
(
    name   VARCHAR(255) PRIMARY KEY,
    active BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS map_configuration_item
(
    id                   BIGSERIAL PRIMARY KEY,
    map_configuration_id BIGINT  NOT NULL REFERENCES map_configuration (id),
    item_id              BIGINT  NOT NULL REFERENCES item (id),
    position_x           INTEGER NOT NULL,
    position_y           INTEGER NOT NULL
);

CREATE TYPE rule_theme AS ENUM ('normal', 'conditional', 'loop');

CREATE TABLE IF NOT EXISTS map_configuration_rule
(
    map_configuration_id BIGINT       NOT NULL REFERENCES map_configuration (id),
    rule                 VARCHAR(255) NOT NULL REFERENCES rule (name),
    theme                rule_theme   NOT NULL,
    parameters           INTEGER[]    NOT NULL
);

CREATE TABLE IF NOT EXISTS map_configuration_door
(
    id                   BIGSERIAL PRIMARY KEY,
    map_configuration_id BIGINT        NOT NULL REFERENCES map_configuration (id),
    door_id              BIGINT        NOT NULL REFERENCES door (id),
    position_x           INTEGER       NOT NULL,
    position_y           INTEGER       NOT NULL,
    door_direction       map_direction NOT NULL
);

CREATE TABLE IF NOT EXISTS map_configuration_for_player
(
    id                   BIGSERIAL PRIMARY KEY,
    map_configuration_id BIGINT       NOT NULL REFERENCES map_configuration (id),
    player_id            VARCHAR(255) NOT NULL REFERENCES player (player_id),
    is_pass              BOOLEAN      NOT NULL
);

CREATE TABLE IF NOT EXISTS game_session
(
    id                   BIGSERIAL PRIMARY KEY,
    player_id            VARCHAR(255) NOT NULL REFERENCES player (player_id),
    map_configuration_id BIGINT       NOT NULL REFERENCES map_configuration (id),
    start_datetime       TIMESTAMPTZ  NOT NULL DEFAULT (now()),
    end_datetime         TIMESTAMPTZ
);

CREATE TYPE medal_type AS ENUM ('gold', 'silver', 'bronze');

CREATE TABLE IF NOT EXISTS play_history
(
    id                BIGSERIAL PRIMARY KEY,
    game_session_id   BIGINT      NOT NULL REFERENCES game_session (id),
    action_step       INTEGER     NOT NULL,
    number_of_command INTEGER     NOT NULL,
    is_finited        BOOLEAN     NOT NULL,
    is_completed      BOOLEAN     NOT NULL,
    command_medal     medal_type,
    action_medal      medal_type,
    submit_datetime   TIMESTAMPTZ NOT NULL DEFAULT (now())
);

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

CREATE TYPE command_edge_type AS ENUM ('conditional_branch', 'main_branch');

CREATE TABLE IF NOT EXISTS command_edge
(
    source_node_id      BIGINT            NOT NULL REFERENCES command_node (id),
    destination_node_id BIGINT            NOT NULL REFERENCES command_node (id),
    type                command_edge_type NOT NULL
);

CREATE TABLE IF NOT EXISTS play_history_rule
(
    play_history_id BIGINT       NOT NULL REFERENCES play_history (id),
    rule            VARCHAR(255) NOT NULL REFERENCES rule (name),
    value           INTEGER      NOT NULL,
    is_pass         BOOLEAN      NOT NULL
);