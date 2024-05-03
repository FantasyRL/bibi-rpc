package rpc

import (
	"bibi/kitex_gen/chat/chathandler"
	"bibi/kitex_gen/follow/followhandler"
	"bibi/kitex_gen/interaction/interactionhandler"
	"bibi/kitex_gen/user/userhandler"
	"bibi/kitex_gen/video/videohandler"
)

var (
	userClient        userhandler.Client
	followClient      followhandler.Client
	interactionClient interactionhandler.Client
	chatClient        chathandler.Client
	videoClient       videohandler.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
	InitFollowRPC()
	InitInteractionRPC()
	InitChatRPC()

}
