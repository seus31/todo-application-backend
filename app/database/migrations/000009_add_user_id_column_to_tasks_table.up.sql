ALTER TABLE tasks ADD COLUMN user_id INTEGER NOT NULL REFERENCES users(id);
