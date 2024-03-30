package main

import (
	user "kitex_gen/user/userhandler"
	"log"
)

func main() {
	svr := user.NewServer(new(UserHandlerImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}