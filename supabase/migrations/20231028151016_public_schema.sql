CREATE TABLE setting_types(
    setting_type_id serial PRIMARY KEY,
    setting_type character varying(20) NOT NULL UNIQUE
);

CREATE TABLE settings(
    setting_id serial PRIMARY KEY,
    setting_type_id int NOT NULL,
    setting_key character varying(100) NOT NULL UNIQUE,
    setting_value text NOT NULL
);
 

-- Alter tables within the "users" schema
ALTER TABLE settings
    ADD FOREIGN KEY (setting_type_id) REFERENCES setting_types(setting_type_id);
 
