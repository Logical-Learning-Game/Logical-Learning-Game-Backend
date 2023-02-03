-- name: GetItemFromMapConfigurationIDs :many
SELECT map_configuration_id, item.id AS item_id, item.name, item.type, map_item.position_x, map_item.position_y
FROM map_configuration_item AS map_item
         INNER JOIN item ON item.id = map_item.item_id
WHERE map_configuration_id = ANY (sqlc.arg(map_configuration_ids):: BIGINT []);