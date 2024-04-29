CREATE TABLE IF NOT EXISTS nothifications(
    id INTEGER NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL, 
    post_id INTEGER NOT NULL, 
	type TEXT NOT NULL CHECK (type IN ('like', 'dislike', 'comment')),
    created_user_id INTEGER NOT NULL,
	message TEXT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT (datetime('now')),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
    FOREIGN KEY (created_user_id) REFERENCES users(id) ON DELETE CASCADE

) 