package eventbus

import "context"

type Message struct {
	UUID     string
	Metadata map[string]string
	Payload  []byte
}

type Publisher interface {
	Publish(topic string, messages ...*Message) error
	Close() error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan *Message, error)
	Close() error
}
