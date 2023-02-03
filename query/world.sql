-- name: ListWorldFromMapConfigurationIDs :many
SELECT id AS world_id, name
FROM world
WHERE id = ANY (sqlc.arg(map_configuration_ids):: BIGINT []);