CREATE TABLE IF NOT EXISTS comments_likes (
    user_id INTEGER NOT NULL,
    comment_id INTEGER NOT NULL,
    is_like BOOLEAN DEFAULT true NOT NULL,
    up INTEGER NOT NULL,
    down INTEGER NOT NULL,
    FOREIGN KEY (post_id) REFERENCES comments(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);