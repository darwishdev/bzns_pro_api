-- name: EntitiesList :many
SELECT
    entity_id,
    entity_name,
    entity_code,
    created_at,
    updated_at,
    deleted_at
FROM
    entities_schema.entities
WHERE
    entity_type_id = IsZero($1, entity_type_id);

-- name: EntitiesInputList :many
SELECT
    entity_id,
    entity_name
FROM
    entities_schema.entities
WHERE
    deleted_at IS NULL
    AND entity_type_id = IsZero($1, entity_type_id);

-- name: EntityDeleteRestore :exec
UPDATE
    entities_schema.entities
SET
    deleted_at = CASE WHEN deleted_at IS NULL THEN
        now()
    ELSE
        NULL
    END
WHERE
    entity_id = ANY (sqlc.arg('entity_ids')::int[]);

-- name: EntityCreate :one
INSERT INTO entities_schema.entities(entity_name, entity_code, entity_type_id)
    VALUES ($1, $2, $3)
RETURNING
    *;



-- name: EntityUpdate :one
UPDATE
    entities_schema.entities
SET
    entity_name = $2,
    entity_code = $3
WHERE
    entity_id = $1
RETURNING
    *;

-- name: EntityFindForUpdate :one
SELECT
    entity_id,
    entity_name,
    entity_code
FROM
    entities_schema.entities
WHERE
    entity_id = $1;

-- -- name: BranchFindForPos :one
-- SELECT
--     b.entity_id,
--     b.entity_name,
--     b.entity_code,
--     b.created_at,
--     b.deleted_at,
--     b.halls
-- FROM
--     entities_schema.branches_view b
-- WHERE
--     b.entity_id = $1;

