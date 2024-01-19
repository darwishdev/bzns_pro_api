
-- name: WarehouseFind :one
select 
	entity_id,
	entity_name,
	entity_code,
	created_at,
	updated_at,
	deleted_at,
	is_private,
	items
from entities_schema.warehouses_view
where entity_id = $1;