CREATE TABLE IF NOT EXISTS wants(
		user_id INTEGER NOT NULL,
		user_name TEXT NOT NULL,
        status TEXT NOT NULL DEFAULT 0 CHECK (status IN (0, 1, -1)),
		-- created_at TIMESTAMP DEFAULT (datetime('now')), -- MB REMOVE
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);