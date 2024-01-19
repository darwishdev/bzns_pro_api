-- name: SettingsUpdate :one
SELECT
FROM
    SettingsBulkUpdate(sqlc.arg('keys')::text[], sqlc.arg('values')::text[]);

-- name: SettingsFindForUpdate :many
SELECT
    setting_value,
    setting_key,
    setting_type
FROM
    settings s
    JOIN setting_types t ON t.setting_type_id = s.setting_type_id;

