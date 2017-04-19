package pubsub

type Message struct {
	Topic   string
	Payload []byte
}

type HandlerFunc func(Message) error
