package rpc

import "bibi/kitex_gen/user/userhandler"

var (
	userClient userhandler.Client
	//followClient      followservice.Client
	//interactionClient interactionservice.Client
	//chatClient        messageservice.Client
	//videoClient       videoservice.Client
)

func Init() {
	InitUserRPC()
	//InitFollowRPC()
	//InitInteractionRPC()
	//InitChatRPC()
	//InitVideoRPC()
}
