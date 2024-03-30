package rpc

import (
	"bibi/kitex_gen/user/userhandler"
	"bibi/kitex_gen/video/videohandler"
)

var (
	userClient userhandler.Client
	//followClient      followservice.Client
	//interactionClient interactionservice.Client
	//chatClient        messageservice.Client
	videoClient videohandler.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
	//InitFollowRPC()
	//InitInteractionRPC()
	//InitChatRPC()

}
