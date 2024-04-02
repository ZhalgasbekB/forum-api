CREATE TABLE IF NOT EXISTS comments_likes (
    user_id INTEGER NOT NULL,
    comment_id INTEGER NOT NULL,
    status BOOLEAN DEFAULT false NOT NULL,
    like_code INTEGER NOT NULL, 
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);