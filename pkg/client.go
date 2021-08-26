package pkg

import (
	"net"
	"sync"
	"time"
)

type Client struct {
	Addr    string
	Timeout time.Duration

	mu   sync.Mutex
	conn net.Conn // 構造体にconnをもたせてコネクションプール
}

func NewClient(addr string, timeout time.Duration) *Client {
	return &Client{
		Addr:    addr,
		Timeout: timeout,
	}
}

func (c *Client) Ping(b []byte) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.connect(); err != nil {
		return "", err
	}

	if c.Timeout > 0 {
		if err := c.conn.SetDeadline(time.Now().Add(c.Timeout)); err != nil {
			return "", err
		}
	}
	_, err := c.conn.Write(b)
	if err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	if c.Timeout > 0 {
		if err = c.conn.SetDeadline(time.Now().Add(c.Timeout)); err != nil {
			return "", err
		}
	}
	n, err := c.conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.close()
}

func (c *Client) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.connect()
}

func (c *Client) close() error {
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) connect() error {
	if c.conn == nil {
		conn, err := net.DialTimeout("tcp", c.Addr, c.Timeout)
		if err != nil {
			return err
		}
		c.conn = conn
	}
	return nil
}
