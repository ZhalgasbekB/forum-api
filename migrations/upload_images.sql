CREATE TABLE images (
    id SERIAL PRIMARY KEY,
    path VARCHAR,
    post_id INT,
    comment_id INT,
    is_active INT DEFAULT 1
);
