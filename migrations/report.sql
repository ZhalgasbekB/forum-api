CREATE TABLE reports (
    report_id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER,
    comment_id INTEGER,
    moderator INTEGER NOT NULL,
    admin INTEGER,
    status TEXT NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'closed', 'ignored')),
    reason TEXT NOT NULL,
    admin_response TEXT,
    created_at TIMESTAMP DEFAULT (datetime('now')),
    updated_at TIMESTAMP DEFAULT (datetime('now')),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE SET NULL,
    FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE SET NULL,
    FOREIGN KEY (moderator) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (admin) REFERENCES users(id) ON DELETE SET NULL
);