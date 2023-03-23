CREATE TYPE map_direction AS ENUM
    (
        'up',
        'left',
        'down',
        'right'
        );

CREATE TYPE item_type AS ENUM ('key_a', 'key_b', 'key_c');

CREATE TYPE door_type AS ENUM ('door_no_key', 'door_a', 'door_b', 'door_c');

CREATE TYPE map_difficulty AS ENUM ('easy', 'medium', 'hard');

CREATE TYPE rule_theme AS ENUM ('normal', 'conditional', 'loop');

CREATE TYPE medal_type AS ENUM ('gold', 'silver', 'bronze', 'none');

CREATE TYPE command_node_type AS ENUM
    (
        'start',
        'conditional_a',
        'conditional_b',
        'conditional_c',
        'conditional_d',
        'conditional_e',
        'forward',
        'left',
        'back',
        'right'
        );

CREATE TYPE command_edge_type AS ENUM ('conditional_branch', 'main_branch');