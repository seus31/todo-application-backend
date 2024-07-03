CREATE TABLE tasks (
    task_id integer PRIMARY KEY,
    task_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);