CREATE OR REPLACE FUNCTION IsZero(in_value ANYELEMENT, in_target_value ANYELEMENT)
    RETURNS int
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF in_value = 0 THEN
        RETURN in_target_value;
    ELSE
        RETURN in_value;
    END IF;
END
$$;

CREATE OR REPLACE FUNCTION settings_bulk_create(keys text[], vals text[])
    RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Create a temporary table to hold the new values
    CREATE TEMP TABLE temp_settings AS
    SELECT
        unnest($1) AS setting_key,
        unnest($2) AS setting_value;
    -- Update the main table based on the temporary table
    UPDATE
        settings AS s
    SET
        setting_value = t.setting_value
    FROM
        temp_settings AS t
    WHERE
        s.setting_key = t.setting_key;
    -- Drop the temporary table
    DROP TABLE temp_settings;
END
$$;

CREATE OR REPLACE FUNCTION status_parser(closed_at timestamp, committed_at timestamp, canceled_at timestamp)
    RETURNS varchar (
        200
)
    AS $$
BEGIN
    IF closed_at IS NULL THEN
        RETURN 'open';
    ELSIF closed_at IS NOT NULL
            AND committed_at IS NULL
            AND canceled_at IS NULL THEN
            RETURN 'closed';
    ELSIF canceled_at IS NOT NULL THEN
        RETURN 'canceled';
    ELSE
        RETURN 'committed';
    END IF;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION random_between(min_value int, max_value int)
    RETURNS int
    AS $$
DECLARE
    v_random_value int;
BEGIN
    SELECT
        floor(random() *(max_value - min_value + 1)) + min_value INTO v_random_value;
    RETURN v_random_value;
END;
$$
LANGUAGE plpgsql;

