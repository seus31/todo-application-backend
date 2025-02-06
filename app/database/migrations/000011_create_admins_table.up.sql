CREATE TABLE admins (
                       id SERIAL NOT NULL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL UNIQUE,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP NULL,
                       updated_at TIMESTAMP NULL,
                       deleted_at TIMESTAMP NULL
);
