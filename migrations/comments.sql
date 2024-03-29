CREATE TABLE IF NOT EXISTS comments(
		id integer not null primary key,
		post_id INTEGER NOT NULL, 
		user_id INTEGER NOT NULL, 
		description TEXT NOT NULL,
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);