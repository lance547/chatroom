package main

import (
	"chatroom/global"
	"chatroom/server"
	"fmt"
	"log"
	"net/http"
)

func init() {
	global.Init()
}

func main() {
	var (
		addr   = ":2022"
		banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |
	
	聊天室
`
	)
	fmt.Printf(banner+"start on %s\n", addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
