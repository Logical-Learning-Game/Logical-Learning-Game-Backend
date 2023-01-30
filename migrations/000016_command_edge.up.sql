CREATE TYPE command_edge_type AS ENUM ('conditional_branch', 'main_branch');

CREATE TABLE IF NOT EXISTS command_edge
(
    source_node_id      BIGINT            NOT NULL REFERENCES command_node (id),
    destination_node_id BIGINT            NOT NULL REFERENCES command_node (id),
    type                command_edge_type NOT NULL
);