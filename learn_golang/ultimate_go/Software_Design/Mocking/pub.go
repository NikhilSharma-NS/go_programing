package pub

import "fmt"

type PubSub struct {
	host string
}

func New(host string) *PubSub {
	return &PubSub{
		host: host,
	}
}

func (ps *PubSub) Publish() {
	fmt.Println("Published")
}

func (ps *PubSub) Subscrib() {
	fmt.Println("Subscribed")
}
