-- name: GetRuleFromMapConfigurationIDs :many
SELECT map_configuration_id, rule, theme, parameters
FROM map_configuration_rule
WHERE map_configuration_id = ANY (sqlc.arg(map_configuration_ids):: BIGINT []);