package service

import (
	"gitea.com/lzhuk/forum/internal/service/admin"
	"gitea.com/lzhuk/forum/internal/service/comment"
	"gitea.com/lzhuk/forum/internal/service/nothification"
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
	Admin          admin.IAdminService
	Nothification  nothification.INotificationService //
}

func NewService(userService user.IUserService,
	sessionService user.ISessionService,
	postsService post.IPostsService,
	commentService comment.ICommentService,
	LikePosts post.ILikePostService,
	LikeComments comment.ILikeCommentService,
	Admins admin.IAdminService,
	Nothification1 nothification.INotificationService) Service {
	return Service{
		UserService:    userService,
		SessionService: sessionService,
		PostsService:   postsService,
		CommentService: commentService,
		LikePosts:      LikePosts,
		LikeComments:   LikeComments,
		Admin:          Admins,
		Nothification:  Nothification1,
	}
}
