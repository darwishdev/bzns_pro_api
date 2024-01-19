-- name: AccountsList :many
SELECT
    account_id,
    account_name,
    account_image,
    account_email,
    account_phone,
    created_at,
    updated_at,
    deleted_at
FROM
    accounts_schema.accounts
WHERE
    account_type_id = IsZero($1, account_type_id);

-- name: AccountsInputList :many
SELECT
    account_id,
    account_name
FROM
    accounts_schema.accounts
WHERE
    deleted_at IS NULL
    AND account_type_id = IsZero($1, account_type_id)
    AND is_private IS FALSE;

-- name: AccountDeleteRestore :exec
UPDATE
    accounts_schema.accounts
SET
    deleted_at = CASE WHEN deleted_at IS NULL THEN
        now()
    ELSE
        NULL
    END
WHERE
    account_id = ANY (sqlc.arg('account_ids')::int[]);

-- name: AccountCreate :one
INSERT INTO accounts_schema.accounts(account_name,account_code, account_image, account_email, account_phone, account_type_id)
    VALUES ($1, $2, $3, $4, $5 , $6)
RETURNING
    *;

-- name: AccountUpdate :one
UPDATE
    accounts_schema.accounts
SET
    account_name = $2,
    account_image = $3,
    account_email = $4,
    account_phone = $5,
    account_code = $6
WHERE
    account_id = $1
RETURNING
    *;

-- name: AccountFindForUpdate :one
SELECT
    account_id,
    account_name,
    account_image,
    account_email,
    account_phone
FROM
    accounts_schema.accounts
WHERE
    account_id = $1;

-- name: AccountFind :one
SELECT
    account_id,
    account_name,
    account_image,
    account_email,
    account_phone,
    created_at,
    updated_at,
    deleted_at
FROM
    accounts_schema.accounts
WHERE
    account_id = $1;

