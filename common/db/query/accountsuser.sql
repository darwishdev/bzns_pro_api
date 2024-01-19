-- name: UserFindByEmailOrCode :one
SELECT
        user_id,
        user_name,
<<<<<<< HEAD
         
=======
        account_id,
        entity_id,
        session_id,
        device_id,
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
        user_image,
        account_code,
        user_email,
        user_phone,
        user_password,
        created_at,
        updated_at,
        deleted_at,
        permission_groups,
        side_bar
    FROM
        accounts_schema.users_view
    WHERE
        user_email = $1
<<<<<<< HEAD
        OR user_phone = $1) user_row;
=======
        OR account_code = $1;
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212

-- name: UserPermissionsList :many
SELECT
    permission_group::varchar(200),
    permissions::jsonb
FROM
    UserPermissionsList($1) AS permission_row;

-- name: UserFind :one
SELECT
    Row_to_json(user_row) user_record
FROM (
    SELECT
        user_id,
        user_name,
        user_image,
        user_email,
<<<<<<< HEAD
=======
        account_id,
        entity_id,
        session_id,
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
        user_phone,
        user_password,
        created_at,
        updated_at,
        deleted_at
    FROM
        accounts_schema.users_view
    WHERE
        user_id = $1) user_row;

-- name: UserResetPassword :exec
UPDATE
    accounts_schema.users
SET
    user_password = $2
WHERE
    user_email = $1;

-- name: UsersList :many
SELECT
    Row_to_json(user_record) user_row
FROM (
    SELECT
        user_id,
        user_name,
        user_image,
        user_email,
        user_phone,
        created_at,
        deleted_at
    FROM
        accounts_schema.users_view
    GROUP BY
        user_id,
        user_name,
        user_image,
        user_email,
        user_phone,
        created_at,
        deleted_at) user_record;

-- name: UserDeleteRestore :exec
UPDATE
    accounts_schema.users
SET
    deleted_at = CASE WHEN deleted_at IS NULL THEN
        now()
    ELSE
        NULL
    END
WHERE
    user_id = ANY (sqlc.arg('user_ids')::int[]);

-- name: UserCreate :one
INSERT INTO accounts_schema.users(user_name, user_image, user_email, user_phone, user_password)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    *;

-- name: UserUpdate :one
UPDATE
    accounts_schema.users
SET
    user_name = $2,
    user_image = $3,
    user_email = $4,
    user_phone = $5,
    user_password = sqlc.arg('user_password')
WHERE
    user_id = $1
RETURNING
    *;

-- name: UserPermissionsBulkCreate :copyfrom
INSERT INTO accounts_schema.user_permissions(user_id, permission_id)
    VALUES ($1, $2);

-- name: UserRolesBulkCreate :copyfrom
INSERT INTO accounts_schema.user_roles(user_id, role_id)
    VALUES ($1, $2);

-- name: UserPermissionsClear :exec
DELETE FROM accounts_schema.user_permissions
WHERE user_id = $1;

-- name: UserRolesClear :exec
DELETE FROM accounts_schema.user_roles
WHERE user_id = $1;

-- name: UserFindForUpdate :one
WITH userPermissions AS (
    SELECT
        up.user_id,
        up.permission_id
    FROM
        accounts_schema.users u
        JOIN accounts_schema.user_permissions up ON up.user_id = u.user_id
    WHERE
        up.user_id = $1
),
userRoles AS (
    SELECT
        ur.user_id,
        ur.role_id
    FROM
        accounts_schema.users u
        JOIN accounts_schema.user_roles ur ON ur.user_id = u.user_id
    WHERE
        ur.user_id = $1
)
SELECT
    Row_to_json(user_record) user_row
FROM (
    SELECT
        u.user_id,
        user_name,
        user_image,
        user_email,
        user_phone,
        user_password,
        created_at,
        updated_at,
(
            SELECT
                Array_agg(ur.role_id) roles
            FROM
                userRoles ur
            WHERE
                ur.user_id = u.user_id),(
                SELECT
                    Array_agg(up.permission_id) permissions
                FROM
                    userPermissions up
                WHERE
                    up.user_id = u.user_id)
            FROM
                accounts_schema.users u
            WHERE
                u.user_id = $1
            GROUP BY
                u.user_id,
                user_name,
                user_image,
                user_email,
                user_phone,
                user_password,
                created_at,
                updated_at) user_record;

