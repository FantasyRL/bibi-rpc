package service

import (
	"bibi/kitex_gen/interaction"
	"bibi/rpc/interaction/dal/db"
	"context"
)

type InteractionService struct {
	ctx context.Context
}

func NewInteractionService(ctx context.Context) *InteractionService {
	return &InteractionService{ctx: ctx}
}

func BuildCommentResp(comment *db.Comment) *interaction.Comment {
	return &interaction.Comment{
		Id:       comment.ID,
		VideoId:  comment.VideoID,
		ParentId: &comment.ParentID,
		//User:        BuildUserResp(commenter),
		Content:     comment.Content,
		PublishTime: comment.CreatedAt.Format("2006-01-02 15:01:04"),
	}
}

func BuildCommentsResp(comments []db.Comment) (commentsResp []*interaction.Comment) {
	for _, comment := range comments {
		commentsResp = append(commentsResp, BuildCommentResp(&comment))
	}
	return
}
