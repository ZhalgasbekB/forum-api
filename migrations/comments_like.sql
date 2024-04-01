CREATE TABLE IF NOT EXISTS posts_votes (
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    vote INTEGER NOT NULL
);