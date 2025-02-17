/*
File Name: create_query_functions.sql
Abstract: This file contains functions that provide a convenient way
to interact with the database as they encapsulate common operations
and help reduce the complexity and length of the queries when using
the database driver.

Author: Awais khan <contact@awaiskhan.com.pk>
Created: 07/02/2025
Last Updated: 07/02/2025
*/

-- ======== SCHEMAS ========
CREATE SCHEMA vector;

-- ======== TABLES ========
CREATE TABLE vector.databases
(
    -- ======== KEYS ========
    id            SERIAL        not null
            primary key,
    name          varchar(100)  not null,
    description   text          not null,
    created_at    date          not null,
);

ALTER TABLE vector.databases
    owner to api;

-- Add new column to the table vector.databases "connection_string"

ALTER TABLE vector.databases
    ADD COLUMN connection_string varchar(100) not null;

-- ======== FUNCTIONS ========
-- Function to create a new database

CREATE OR REPLACE FUNCTION vector.create_database(
    name varchar(100),
    description text,
    connection_string varchar(100)
)

RETURNS void AS
$$
BEGIN
    INSERT INTO vector.databases (name, description, created_at, connection_string)
    VALUES (name, description, now(), connection_string);
END;

$$ LANGUAGE plpgsql;

-- Function to get all databases

CREATE OR REPLACE FUNCTION vector.get_databases()
RETURNS SETOF vector.databases AS
$$


BEGIN
    RETURN QUERY SELECT * FROM vector.databases;
END;


$$ LANGUAGE plpgsql;

-- Function to get a database by id

CREATE OR REPLACE FUNCTION vector.get_database_by_id(
    id int
)

RETURNS SETOF vector.databases AS
$$


BEGIN
    RETURN QUERY SELECT * FROM vector.databases WHERE id = id;
END;


$$ LANGUAGE plpgsql;

-- Function to update a database

CREATE OR REPLACE FUNCTION vector.update_database(
    id int,
    name varchar(100),
    description text,
    connection_string varchar(100)
)

RETURNS void AS
$$
BEGIN
    UPDATE vector.databases
    SET name = name, description = description, connection_string = connection_string
    WHERE id = id;
END;

$$ LANGUAGE plpgsql;

-- Function to delete a database

CREATE OR REPLACE FUNCTION vector.delete_database(
    id int
)

RETURNS void AS
$$
BEGIN
    DELETE FROM vector.databases WHERE id = id;
END;



