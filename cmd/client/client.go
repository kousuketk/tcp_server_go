package main

import (
	"fmt"

	"github.com/kousuketk/tcp_server_go/pkg"
)

func main() {
	c := pkg.NewClient(":5000")

	resp, err := c.Ping([]byte("hello"))
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
