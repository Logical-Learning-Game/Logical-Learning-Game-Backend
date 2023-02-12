-- name: GetRuleFromMapConfigurationIDs :many
SELECT id, map_configuration_id, rule, rule_order, theme, parameters
FROM map_configuration_rule
WHERE map_configuration_id = ANY (sqlc.arg(map_configuration_ids):: BIGINT []);