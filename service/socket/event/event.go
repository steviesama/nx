package event

import "github.com/steviesama/nx/rand"

func init() {
	guid := rand.Guid(true)
}

type Publisher interface {
	PublishEvent(Event) error
}

type Subscriber interface {
	Subscribe()
}
