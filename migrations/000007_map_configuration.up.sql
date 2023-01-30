CREATE TYPE map_difficulty AS ENUM ('easy', 'medium', 'hard');

CREATE TABLE IF NOT EXISTS map_configuration
(
    id                            BIGSERIAL PRIMARY KEY,
    world_id                      BIGINT         NOT NULL REFERENCES world (id),
    map_id                        BIGINT         NOT NULL REFERENCES map (id),
    config_name                   VARCHAR(255)   NOT NULL,
    map_image_path                VARCHAR(255),
    difficulty                    map_difficulty NOT NULL,
    star_requirement              INT            NOT NULL,
    least_solvable_command_gold   INT            NOT NULL,
    least_solvable_command_silver INT            NOT NULL,
    least_solvable_command_bronze INT            NOT NULL,
    least_solvable_action_gold    INT            NOT NULL,
    least_solvable_action_silver  INT            NOT NULL,
    least_solvable_action_bronze  INT            NOT NULL,
    created_at                    TIMESTAMPTZ    NOT NULL DEFAULT (now())
);