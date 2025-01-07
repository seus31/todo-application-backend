CREATE TABLE task_dependencies (
    id SERIAL NOT NULL PRIMARY KEY,
    task_id INTEGER NOT NULL REFERENCES tasks(id),
    dependent_task_id INTEGER NOT NULL REFERENCES tasks(id),
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);