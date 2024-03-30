package main

import (
	video "bibi/kitex_gen/video/videohandler"
	"log"
)

func main() {
	svr := video.NewServer(new(VideoHandlerImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
