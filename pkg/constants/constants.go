package constants

import "time"

const (
	// service name
	APIServiceName         = "api"
	UserServiceName        = "user"
	InteractionServiceName = "interaction"
	FollowServiceName      = "follow"
	ChatServiceName        = "chat"
	VideoServiceName       = "video"

	// db table name
	UserTableName        = "user"
	ChatTableName        = "message"
	CommentTableName     = "comment"
	CommentLikeTableName = "comment_like"
	FavoriteTableName    = "like"
	FollowTableName      = "follow"
	VideoTableName       = "video"

	// limit
	MaxConnections     = 1000
	MaxQPS             = 100
	MaxRequestBodySize = 114514 * 1024 * 1024
	MaxIdleConns       = 20
	MaxGoroutines      = 10
	MaxOpenConns       = 100
	ConnMaxLifetime    = 10 * time.Second

	// RPC
	MuxConnection  = 1
	RPCTimeout     = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond
)

const (
	// page
	PageNum  = 1
	PageSize = 10

	// interaction type
	AddComment    = 1
	DeleteComment = 0
	Like          = 1
	Dislike       = 0

	// follow type
	FollowAction   = 1
	UnFollowAction = 0
)
