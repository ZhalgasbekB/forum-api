CREATE TABLE IF NOT EXISTS posts (
    id INTEGER NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    category_name TEXT NOT NULL,
    title TEXT NOT NULL,
    discription TEXT,
    create_at DATETIME NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
