package event

type Publisher interface {
	PublishEvent() error
}

type Subscriber interface {
}

type Event struct {
}
