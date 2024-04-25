run: migrate
	go run cmd/main.go
migrate:
	sqlite3 forum.sqlite3 < migrations/users.sql
	sqlite3 forum.sqlite3 < migrations/sessions.sql
	sqlite3 forum.sqlite3 < migrations/posts.sql
	sqlite3 forum.sqlite3 < migrations/posts_like.sql
	sqlite3 forum.sqlite3 < migrations/comments.sql
	sqlite3 forum.sqlite3 < migrations/comments_like.sql
	sqlite3 forum.sqlite3 < migrations/report.sql



