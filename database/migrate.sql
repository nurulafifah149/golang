CREATE DATABASE project;

CREATE TABLE
    books(
        id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        title VARCHAR(255),
        author VARCHAR(255),
        dsc VARCHAR(255),
        created_at timestamptz DEFAULT now(),
        updated_at timestamptz DEFAULT now(),
        deleted_at timestamptz
    );