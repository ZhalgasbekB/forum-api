CREATE TABLE IF NOT EXISTS users(
		id INTEGER NOT NULL PRIMARY KEY,
		name VARCHAR(100) UNIQUE,
		email VARCHAR(100) UNIQUE, 
		password VARCHAR(100),
        role VARCHAR(30),
		created_at TIMESTAMP
);