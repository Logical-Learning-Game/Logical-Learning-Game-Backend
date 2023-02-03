-- name: GetDoorFromMapConfigurationIDs :many
SELECT map_configuration_id,
       door.id AS door_id,
       door.name,
       door.type,
       map_door.position_x,
       map_door.position_y,
       map_door.door_direction
FROM map_configuration_door AS map_door
         INNER JOIN door
                    ON door.id = map_door.door_id
WHERE map_configuration_id = ANY (sqlc.arg(map_configuration_ids):: BIGINT []);