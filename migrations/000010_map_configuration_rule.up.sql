CREATE TABLE IF NOT EXISTS map_configuration_rule
(
    map_configuration_id BIGINT       NOT NULL REFERENCES map_configuration (id),
    rule                 VARCHAR(255) NOT NULL REFERENCES rule (name),
    parameters           INTEGER
);