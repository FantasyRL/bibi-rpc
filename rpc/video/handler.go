package main

import (
	video "bibi/kitex_gen/video"
	"context"
)

// VideoHandlerImpl implements the last service interface defined in the IDL.
type VideoHandlerImpl struct{}

// PutVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) PutVideo(ctx context.Context, req *video.PutVideoRequest) (resp *video.PutVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// ListVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) ListVideo(ctx context.Context, req *video.ListUserVideoRequest) (resp *video.ListUserVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) SearchVideo(ctx context.Context, req *video.SearchVideoRequest) (resp *video.SearchVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// HotVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) HotVideo(ctx context.Context, req *video.HotVideoRequest) (resp *video.HotVideoResponse, err error) {
	// TODO: Your code here...
	return
}
