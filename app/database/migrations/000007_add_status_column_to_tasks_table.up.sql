CREATE TYPE task_status AS ENUM ('not_yet_started', 'in_progress', 'completed');
ALTER TABLE tasks ADD COLUMN status task_status NULL;
