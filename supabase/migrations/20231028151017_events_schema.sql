CREATE SCHEMA events_schema;
 
CREATE TABLE events_schema.categories(
    category_id serial PRIMARY KEY,
    category_name character varying(200) NOT NULL UNIQUE,
     created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
); 

CREATE TABLE events_schema.events(
    event_id serial PRIMARY KEY,
    event_name character varying(200) NOT NULL UNIQUE,
    event_location character varying(200) NOT NULL,
    event_location_url character varying(200) NOT NULL,
    constructor_name character varying(200) NOT NULL,
    constructor_title character varying(200) NOT NULL,
    constructor_image character varying(200) NOT NULL,
    event_plan text  NOT NULL,
    event_goals text NOT NULL,
    event_breif text,
    event_description text,
    event_video text,
    event_date date,
    event_start_time time,
    event_end_time time,
     event_hours int,
    category_id int not null,
    event_image character varying(200),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
); 
-- Alter tables within the users schema
ALTER TABLE events_schema.events
    ADD FOREIGN KEY (category_id) REFERENCES events_schema.categories(category_id) ;
 