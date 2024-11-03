CREATE TABLE tasks (
    id SERIAL NOT NULL PRIMARY KEY,
    task_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);