package rpc

import (
	"bibi/kitex_gen/interaction/interactionhandler"
	"bibi/kitex_gen/user/userhandler"
	"bibi/kitex_gen/video/videohandler"
)

var (
	userClient userhandler.Client
	//followClient      followservice.Client
	interactionClient interactionhandler.Client
	//chatClient        messageservice.Client
	videoClient videohandler.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
	//InitFollowRPC()
	InitInteractionRPC()
	//InitChatRPC()

}
