
-- name: BranchCreate :exec
select entities_schema.branch_create($1,$2,$3,$4,$5,$6,$7,$8 , $9 , $10 , $11 ,  sqlc.arg('in_hall_names')::varchar[], sqlc.arg('in_hall_codes')::varchar[], sqlc.arg('in_hall_tables_count')::int[]);
-- name: BranchFind :one
select 
	entity_id,
	entity_name,
	entity_code,
	created_at,
	updated_at,
	deleted_at,
	is_private,
	settings,
	halls,
	committed_orders,
	active_orders
	 from entities_schema.branches_view
where entity_id = $1;