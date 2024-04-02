package service

import (
	"gitea.com/lzhuk/forum/internal/service/comment"
	"gitea.com/lzhuk/forum/internal/service/post"
	"gitea.com/lzhuk/forum/internal/service/user"
)

type Service struct {
	UserService    user.IUserService
	SessionService user.ISessionService
	PostsService   post.IPostsService
	CommentService comment.ICommentService
	LikePosts      post.ILikePostService
	LikeComments   comment.ILikeCommentService
}

func NewService(userService user.IUserService, sessionService user.ISessionService, postsService post.IPostsService, commentService comment.ICommentService, LikePosts post.ILikePostService, LikeComments comment.ILikeCommentService) Service {
	return Service{UserService: userService, SessionService: sessionService, PostsService: postsService, CommentService: commentService, LikePosts: LikePosts, LikeComments: LikeComments}
}
