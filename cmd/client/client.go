package main

import (
	"fmt"
	"time"

	"github.com/kousuketk/tcp_server_go/pkg"
)

const timeout = 1 * time.Second

func main() {
	c := pkg.NewClient(":5000", timeout)

	resp, err := c.Ping([]byte("hello"))
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
