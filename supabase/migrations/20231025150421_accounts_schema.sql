CREATE SCHEMA accounts_schema;
 
CREATE TABLE accounts_schema.permissions(
    permission_id serial PRIMARY KEY,
    permission_function varchar(200) NOT NULL UNIQUE,
    permission_name varchar(200) NOT NULL,
    permission_description varchar(200),
    permission_group varchar(200) NOT NULL
);

CREATE TABLE accounts_schema.roles(
    role_id serial PRIMARY KEY,
    role_name varchar(200) NOT NULL UNIQUE,
    role_description varchar(200),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp,
    is_private boolean NOT NULL DEFAULT FALSE
);

CREATE TABLE accounts_schema.role_permissions(
    role_permission_id serial PRIMARY KEY,
    role_id int NOT NULL,
    permission_id int NOT NULL
);

CREATE TABLE accounts_schema.users(
    user_id serial PRIMARY KEY,
    user_name varchar(200) NOT NULL,
    user_image varchar(200),
    user_email varchar(200) UNIQUE NOT NULL,
    user_phone varchar(200) UNIQUE,
    user_password varchar(200) NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp,
    is_private boolean NOT NULL DEFAULT FALSE
);

CREATE TABLE accounts_schema.user_roles(
    user_role_id serial PRIMARY KEY,
    user_id int NOT NULL,
    role_id int NOT NULL
);

CREATE TABLE accounts_schema.user_permissions(
    user_permission_id serial PRIMARY KEY,
    user_id int NOT NULL,
    permission_id int NOT NULL
);
   
CREATE TABLE accounts_schema.navigation_bars(
    navigation_bar_id serial PRIMARY KEY,
    menu_key varchar(200) UNIQUE NOT NULL,
    label varchar(200) NOT NULL,
    icon varchar(200),
    "to" varchar(200) UNIQUE,
    parent_id int,
    permission_id int
);

-- Alter tables within the users schema
ALTER TABLE accounts_schema.role_permissions
    ADD FOREIGN KEY (role_id) REFERENCES accounts_schema.roles(role_id),
    ADD FOREIGN KEY (permission_id) REFERENCES accounts_schema.permissions(permission_id);

ALTER TABLE accounts_schema.user_roles
    ADD FOREIGN KEY (user_id) REFERENCES accounts_schema.users(user_id),
    ADD FOREIGN KEY (role_id) REFERENCES accounts_schema.roles(role_id);

ALTER TABLE accounts_schema.user_permissions
    ADD FOREIGN KEY (user_id) REFERENCES accounts_schema.users(user_id),
    ADD FOREIGN KEY (permission_id) REFERENCES accounts_schema.permissions(permission_id);

ALTER TABLE accounts_schema.navigation_bars
    ADD FOREIGN KEY (parent_id) REFERENCES accounts_schema.navigation_bars(navigation_bar_id),
    ADD FOREIGN KEY (permission_id) REFERENCES accounts_schema.permissions(permission_id);
 