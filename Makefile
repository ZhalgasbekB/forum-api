run: migrate
	go run cmd/main.go
migrate:
	sqlite3 forum.sqlite3 < migrations/users.sql
	sqlite3 forum.sqlite3 < migrations/sessions.sql
	sqlite3 forum.sqlite3 < migrations/posts.sql
	sqlite3 forum.sqlite3 < migrations/posts_vote.sql
	sqlite3 forum.sqlite3 < migrations/comments.sql
