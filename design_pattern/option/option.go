package option

import "fmt"

/**
 * 选项设计模式
 */
type Client struct {
	option1 string
	option2 int
}

func (c *Client) String() {
	fmt.Println(c.option1, c.option2)
}

type Option func(*Client)

func NewClient(opts ...Option) Client {
	var c Client
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

func WithOption1(option1 string) Option {
	return func(c *Client) {
		c.option1 = option1
	}
}

func WithOption2(option2 int) Option {
	return func(c *Client) {
		c.option2 = option2
	}
}
