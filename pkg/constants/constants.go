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
	LikeTableName        = "like"
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
	ElasticSearchIndexName = "bibi"

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

	//redis
	VideoLikeSuffix     = ":video_like"
	CommentLikeSuffix   = ":comment_like"
	FollowerCountSuffix = ":follower_counts"
	FriendCountSuffix   = ":friend_counts"
	CommentSuffix       = ":comment"
	FollowerSuffix      = ":follower"
	ReceiveSuffix       = ":receive"

	VideoExpTime   = time.Hour * 1 //到期自动移除k-v
	LikeExpTime    = time.Minute * 10
	CommentExpTime = time.Minute * 10
	FollowExpTime  = time.Minute
	MessageExpTime = time.Hour * 24 * 7

	VideoLikeZset         = "video_likes"
	CommentLikeZset       = "comment_likes"
	VideoCommentCountZset = "video_comment_counts"
	VideoCommentZset      = "video_comments"
	FollowerCountZset     = "follower_counts"
	FollowingCountZset    = "following_counts"
	FriendCountZset       = "friend_counts"
)
