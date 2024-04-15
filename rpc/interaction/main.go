package main

import (
	interaction "bibi/kitex_gen/interaction/interactionhandler"
	"log"
)

func main() {
	svr := interaction.NewServer(new(InteractionHandlerImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
