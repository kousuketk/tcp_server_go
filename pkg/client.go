package pkg

import (
	"net"
	"time"
)

type Client struct {
	Addr    string
	Timeout time.Duration
}

func NewClient(addr string, timeout time.Duration) *Client {
	return &Client{
		Addr:    addr,
		Timeout: timeout,
	}
}

func (c *Client) Ping(b []byte) (string, error) {
	conn, err := net.DialTimeout("tcp", c.Addr, c.Timeout)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if c.Timeout > 0 {
		if err = conn.SetDeadline(time.Now().Add(c.Timeout)); err != nil {
			return "", err
		}
	}
	_, err = conn.Write(b)
	if err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	if c.Timeout > 0 {
		if err = conn.SetDeadline(time.Now().Add(c.Timeout)); err != nil {
			return "", err
		}
	}
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}
