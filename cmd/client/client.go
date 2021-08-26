package client_main

import (
	"fmt"

	"github.com/kousukekt/tcp_server_go/pkg"
)

func main() {
	c := pkg.NewClient(":5000")

	resp, err := c.Ping([]byte("hello"))
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
