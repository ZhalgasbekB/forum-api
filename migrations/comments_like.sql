CREATE TABLE IF NOT EXISTS comments_likes (
    user_id INTEGER NOT NULL,
    comment_id INTEGER NOT NULL,
    status BOOLEAN NOT NULL,
    FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);